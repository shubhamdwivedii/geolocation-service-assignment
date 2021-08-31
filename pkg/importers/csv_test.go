package importers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/geolocation"
	sv "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/service"
	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/storage"
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

	storage, _ := NewMockStorage()

	service := sv.NewService(storage)

	metricChannel, err := importer.Import(path, service)
	require.NoError(t, err)

	metrics := <-metricChannel

	for _, record := range records[1:4] {
		gloc, err := service.GetGeodata(record[0])
		require.NoError(t, err)
		assert := assert.New(t)
		assert.Equal(gloc.IP, record[0], "Parsed IP Should Match.")
		assert.Equal(gloc.CCode, record[1], "Parsed Country Code Should Match.")
		assert.Equal(gloc.Country, record[2], "Parsed Country Should Match.")
		assert.Equal(gloc.City, record[3], "Parsed City Should Match.")
		assert.Equal(fmt.Sprintf("%v", gloc.Latitude), record[4], "Parsed Latitude Should Match.")
		assert.Equal(fmt.Sprintf("%v", gloc.Longitude), record[5], "Parsed Longitude Should Match.")
		assert.Equal(fmt.Sprintf("%v", gloc.MValue), record[6], "Parsed Mystery Value Should Match.")
	}

	// Records after  have intentional errors.
	for _, record := range records[4:] {
		gloc, err := service.GetGeodata(record[0])
		require.Error(t, err)
		assert := assert.New(t)
		assert.Nil(gloc)
	}

	assert.Equal(t, metrics.Imported, 3, "Expected Imported To Be 3.")
	assert.Equal(t, metrics.Rejected, 1, "Expected Rejected To Be 1.")

	// Removing Temp CSV
	err = os.Remove("temp.csv")
	if err != nil {
		log.Println("Error Removing Temp CSV", err.Error())
	}
}

/*********** MOCK DB Storage ***********/
type MockStorage struct {
	db map[string]Geolocation
}

func NewMockStorage() (Storage, error) {
	s := new(MockStorage)
	s.db = make(map[string]Geolocation)
	return s, nil
}

func (s *MockStorage) AddGeodata(gloc Geolocation) error {
	err := ValidateGeolocation(gloc)
	if err != nil {
		return err
	}

	if _, ok := s.db[gloc.IP]; ok {
		return errors.New("Duplicate Entry.")
	} else {
		s.db[gloc.IP] = gloc
	}
	return nil
}

func (s *MockStorage) GetGeodata(ip string) (*Geolocation, error) {
	if gloc, ok := s.db[ip]; !ok {
		return nil, errors.New("Does Not Exits In DB.")
	} else {
		return &gloc, nil
	}
}

func (s *MockStorage) GetAllByCCode(ccode string) ([]*Geolocation, error) {
	var locations []*Geolocation
	for _, gloc := range s.db {
		locations = append(locations, &gloc)
	}
	return locations, nil
}
