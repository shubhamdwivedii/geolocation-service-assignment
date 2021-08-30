package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/geolocation"
	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/storage"
)

func TestHandler(t *testing.T) {
	storage, _ := NewMockStorage()

	rr := httptest.NewRecorder()

	handler := NewHandler(storage)

	oloc := Geolocation{
		IP:        "127.42.24.1",
		CCode:     "IN",
		Country:   "India",
		City:      "Delhi",
		Longitude: -84.87503094689836,
		Latitude:  7.206435933364332,
		MValue:    7823011346,
	}

	storage.AddGeodata(oloc)

	r, err := http.NewRequest(http.MethodGet, "/geodata/127.42.24.1", nil)

	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, r)

	rs := rr.Result()

	body, err := ioutil.ReadAll(rs.Body)

	if err != nil {
		t.Fatal(err)
	}

	var gloc Geolocation
	err = json.Unmarshal(body, &gloc)

	if err != nil {
		t.Fatal(err)
	}

	if gloc.IP != oloc.IP {
		t.Errorf("Expected Response Geolocation IP to be %v but got %v", oloc.IP, gloc.IP)
	}
	if gloc.CCode != oloc.CCode {
		t.Errorf("Expected Response Geolocation CCode to be %v but got %v", oloc.CCode, gloc.CCode)
	}
	if gloc.Country != oloc.Country {
		t.Errorf("Expected Response Geolocation Country to be %v but got %v", oloc.Country, gloc.Country)
	}
	if gloc.City != oloc.City {
		t.Errorf("Expected Response Geolocation City to be %v but got %v", oloc.City, gloc.City)
	}
	if gloc.Longitude != oloc.Longitude {
		t.Errorf("Expected Response Geolocation Longitude to be %v but got %v", oloc.Longitude, gloc.Longitude)
	}
	if gloc.Latitude != oloc.Latitude {
		t.Errorf("Expected Response Geolocation Latitude to be %v but got %v", oloc.Latitude, gloc.Latitude)
	}
	if gloc.MValue != oloc.MValue {
		t.Errorf("Expected Response Geolocation MValue to be %v but got %v", oloc.MValue, gloc.MValue)
	}

}

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
