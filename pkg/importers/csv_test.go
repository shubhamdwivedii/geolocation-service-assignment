package importers_test

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"

	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/importers"
)

func TestCSVParser(t *testing.T) {
	records := [][]string{
		{"ip_address", "country_code", "country", "city", "latitude", "longitude", "mystery_value"},
		{"200.106.141.15", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"},
		{"160.103.7.140", "CZ", "Nicaragua", "New Neva", "-68.31023296602508", "-37.62435199624531", "7301823115"},
		{"70.95.73.73", "TL", "Saudi Arabia", "Gradymouth", "-49.16675918861615", "-86.05920084416894", "2559997162"},
		{"127.01.7.140", "INR", "India", "New Delhi", "-168.31023296602508", "-237.62435199624531", "7301823115"},
	}

	f, err := os.Create("temp.csv")
	defer f.Close()

	if err != nil {
		log.Println("Error Creating File...")
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records) // calls Flush internally

	if err != nil {
		log.Println("Error Writing To Temp CSV...")
	}

	importer := new(CSVImporter)
	pwd, _ := os.Getwd()
	path := pwd + "/temp.csv"
	log.Println(path)
	importChannel, metricChannel, err := importer.Import(path)

	fmt.Println("DO something with", metricChannel)
	// Add test for metric too.

	idx := 1
	for glocation := range importChannel {
		record := records[idx]
		log.Println("Comparing Record:", record, "With Parsed Geolocation", glocation)
		if glocation.IP != record[0] {
			t.Errorf("Parsed IP Not Match.")
		}
		if glocation.CCode != record[1] {
			t.Errorf("Parsed Country Code Not Match.")
		}
		if glocation.Country != record[2] {
			t.Errorf("Parsed Country Not Match.")
		}
		if glocation.City != record[3] {
			t.Errorf("Parsed City Not Match.")
		}
		if fmt.Sprintf("%v", glocation.Latitude) != record[4] {
			t.Errorf("Parsed Latitude Not Match.")
		}
		if fmt.Sprintf("%v", glocation.Longitude) != record[5] {
			t.Errorf("Parsed Longitude Not Match.")
		}
		if fmt.Sprintf("%v", glocation.MValue) != record[6] {
			t.Errorf("Parsed Mystery Value Not Match.")
		}
		idx++
		if idx > len(records) {
			break
		}
	}

	metrics := <-metricChannel

	if metrics.Imported != 3 {
		t.Errorf("Expected Imported To Be 3 But Got %v in Metrics", metrics.Imported)
	}

	if metrics.Rejected != 1 {
		t.Errorf("Expected Rejected To Be 1 But Got %v in Metrics", metrics.Rejected)
	}

	// Removing Temp CSV
	err = os.Remove("temp.csv")
	if err != nil {
		log.Println("Error Removing Temp CSV", err.Error())
	}
}
