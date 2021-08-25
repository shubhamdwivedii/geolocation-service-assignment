package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shubhamdwivedii/geolocation-service-assignment/models"
)

type CSVParser struct {
	metrics models.Metrics
}

func (c *CSVParser) GetMetrics() models.Metrics {
	return c.metrics
}

// Will import Geolocations from CSV file and emit them on a channel
func (*CSVParser) Import(file string) (<-chan models.Geolocation, error) {
	log.Println("Imporint CSV....")

	results := make(chan models.Geolocation)

	csvFile, err := os.Open(file)
	if err != nil {
		log.Println("Error: Reading CSV File:", err.Error())
		return nil, err
	}
	reader := csv.NewReader(csvFile)
	go func() {
		defer close(results)

		header, err := reader.Read()
		if err != nil {
			log.Println("Error:", err.Error())
		}
		log.Println("Headers", header)

		var metrics models.Metrics
		start := time.Now()

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("Error:", err.Error())
			}
			metrics.Total++

			if !validateRecord(record) {
				metrics.Rejected++
				continue
			}
			var gloc models.Geolocation
			gloc.IP = record[0]
			gloc.CCode = record[1]
			gloc.Country = record[2]
			gloc.City = record[3]
			var lterr, lgerr, mverr error
			gloc.Latitude, lterr = strconv.ParseFloat(record[4], 64)
			gloc.Longitude, lgerr = strconv.ParseFloat(record[5], 64)
			gloc.MValue, mverr = strconv.ParseInt(record[6], 10, 64)

			if lterr != nil || lgerr != nil || mverr != nil {
				log.Println("Error Parsing Record..")
				metrics.Rejected++
				continue
			}

			results <- gloc
			metrics.Imported++
		}

		metrics.Duration = time.Since(start)
		// Figure out later how to log this
		log.Println("Metrics", metrics)
	}()
	return results, nil
}

func validateRecord(record []string) bool {
	for _, rec := range record {
		if rec == "" {
			return false
		}
	}
	return true
}
