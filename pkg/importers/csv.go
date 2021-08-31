package importers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/geolocation"
	sv "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/service"
)

type CSVImporter struct {
	metrics Metrics
}

func (c *CSVImporter) GetMetrics() Metrics {
	return c.metrics
}

func NewCSVImporter() Importer {
	importer := new(CSVImporter)
	return importer
}

// Will import Geolocations from CSV file and emit them on a channel
func (*CSVImporter) Import(file string, service sv.Service) (<-chan Metrics, error) {
	log.Println("Importing CSV....")

	mtrchan := make(chan Metrics)

	csvFile, err := os.Open(file)
	if err != nil {
		log.Println("Error: Reading CSV File:", err.Error())
		return nil, err // ERRORSSSSS
	}

	reader := csv.NewReader(csvFile)
	header, err := reader.Read()
	if err != nil {
		log.Println("Error:", err.Error())
	}
	if len(header) != 7 {
		return nil, errors.New("Invalid CSV File.")
	}

	go func() {
		var metrics Metrics
		start := time.Now()

		for {
			record, err := reader.Read()
			if err == io.EOF {
				log.Println("End Of Sample Data Reached...")
				break
			} else if err != nil {
				fmt.Println("Error:", err.Error())
			}
			metrics.Total++

			var gloc *Geolocation
			gloc, err = mapGeolocation(record)
			if err != nil {
				metrics.Rejected++
				continue
			}
			// results <- *gloc
			err = service.AddGeodata(*gloc)
			if err != nil {
				metrics.Rejected++
				continue
			}

			metrics.Imported++
		}
		// Check Why defer above wasn't working...
		metrics.Duration = time.Since(start)
		mtrchan <- metrics
		close(mtrchan)
	}()
	return mtrchan, nil
}

func validateRecord(record []string) bool {
	if len(record) < 6 {
		return false
	}
	for _, rec := range record {
		if rec == "" {
			return false
		}
	}
	return true
}

func mapGeolocation(record []string) (*Geolocation, error) {
	if !validateRecord(record) {
		return nil, errors.New("Invalid Record.")
	}

	var gloc Geolocation
	gloc.IP = record[0]
	gloc.CCode = record[1]
	gloc.Country = record[2]
	gloc.City = record[3]
	var lterr, lgerr, mverr error
	gloc.Latitude, lterr = strconv.ParseFloat(record[4], 64)
	gloc.Longitude, lgerr = strconv.ParseFloat(record[5], 64)
	gloc.MValue, mverr = strconv.ParseInt(record[6], 10, 64)

	if lterr != nil || lgerr != nil || mverr != nil {
		return nil, errors.New("Invalid Record.")
	}
	return &gloc, nil
}
