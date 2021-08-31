package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/geolocation"
	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/storage"
)

func TestNewService(t *testing.T) {
	storage, _ := NewMockStorage()
	service := NewService(storage)

	oloc := Geolocation{
		IP:        "127.42.24.1",
		CCode:     "IN",
		Country:   "India",
		City:      "Delhi",
		Longitude: -84.87503094689836,
		Latitude:  7.206435933364332,
		MValue:    7823011346,
	}

	oloc2 := Geolocation{
		IP:        "192.65.42.10",
		CCode:     "IN",
		Country:   "India",
		City:      "Delhi",
		Longitude: -84.87503094689836,
		Latitude:  7.206435933364332,
		MValue:    7823011346,
	}

	err := service.AddGeodata(oloc)
	require.NoError(t, err)
	err = service.AddGeodata(oloc2)
	require.NoError(t, err)

	gloc, err := service.GetGeodata(oloc.IP)
	require.NoError(t, err)

	assert := assert.New(t)

	assert.Equal(*gloc, oloc, "Expected Both To Be Same.")

	glocs, err := service.GetAllByCCode("IN")
	require.NoError(t, err)

	assert.Equal(len(glocs), 2, "Expected Length To Be 2.")

	for _, gloc := range glocs {
		assert.Equal((*gloc).Country, "INDIA", "Expected Country To Match.")
		assert.Equal((*gloc).City, "DELHI", "Expected City To Match.")
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
