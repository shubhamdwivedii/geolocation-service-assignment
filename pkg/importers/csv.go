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
func (*CSVImporter) Import(file string) (<-chan Geolocation, <-chan Metrics, error) {
	log.Println("Importing CSV....")

	results := make(chan Geolocation)
	mtrchan := make(chan Metrics)

	csvFile, err := os.Open(file)
	if err != nil {
		log.Println("Error: Reading CSV File:", err.Error())
		return nil, nil, err // ERRORSSSSS
	}

	reader := csv.NewReader(csvFile)
	go func() {
		// defer close(results)
		// defer close(mtrchan)
		header, err := reader.Read()
		if err != nil {
			log.Println("Error:", err.Error())
		}
		log.Println("Headers", header)
		// check headers

		var metrics Metrics
		start := time.Now()

		for {
			record, err := reader.Read()
			if err == io.EOF {
				log.Println("END OF FILE REACHED...")
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
			results <- *gloc
			metrics.Imported++
		}
		// Check Why defer above wasn't working...
		close(results)
		metrics.Duration = time.Since(start)
		mtrchan <- metrics
		close(mtrchan)
	}()
	return results, mtrchan, nil
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
