package importers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCSVParser(t *testing.T) {
	records := [][]string{
		{"ip_address", "country_code", "country", "city", "latitude", "longitude", "mystery_value"},
		{"200.106.141.15", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346"},
		{"160.103.7.140", "CZ", "Nicaragua", "New Neva", "-68.31023296602508", "-37.62435199624531", "7301823115"},
		{"70.95.73.73", "TL", "Saudi Arabia", "Gradymouth", "-49.16675918861615", "-86.05920084416894", "2559997162"},
		{"", "INR", "India", "New Delhi", "-168.31023296602508", "-237.62435199624531", "7301823115"},
	}

	f, err := os.Create("temp.csv")
	defer f.Close()
	require.NoError(t, err)

	w := csv.NewWriter(f)
	err = w.WriteAll(records) // calls Flush internally
	require.NoError(t, err)

	importer := new(CSVImporter)
	pwd, _ := os.Getwd()
	path := pwd + "/temp.csv"

	importChannel, metricChannel, err := importer.Import(path)
	require.NoError(t, err)

	idx := 1
	for gloc := range importChannel {
		record := records[idx]
		assert := assert.New(t)

		assert.Equal(gloc.IP, record[0], "Parsed IP Should Match.")

		assert.Equal(gloc.CCode, record[1], "Parsed Country Code Should Match.")

		assert.Equal(gloc.Country, record[2], "Parsed Country Should Match.")

		assert.Equal(gloc.City, record[3], "Parsed City Should Match.")

		assert.Equal(fmt.Sprintf("%v", gloc.Longitude), record[4], "Parsed Longitude Should Match.")

		assert.Equal(fmt.Sprintf("%v", gloc.Latitude), record[5], "Parsed Latitude Should Match.")

		assert.Equal(fmt.Sprintf("%v", gloc.MValue), record[6], "Parsed Mystery Value Should Match.")

		idx++
		if idx > len(records) {
			break
		}
	}

	metrics := <-metricChannel
	assert.Equal(t, metrics.Imported, 3, "Expected Imported To Be 3.")
	assert.Equal(t, metrics.Rejected, 1, "Expected Rejected To Be 1.")

	// Removing Temp CSV
	err = os.Remove("temp.csv")
	if err != nil {
		log.Println("Error Removing Temp CSV", err.Error())
	}
}
