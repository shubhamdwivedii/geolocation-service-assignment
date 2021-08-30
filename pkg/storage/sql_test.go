package storage

import (
	"testing"

	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/geolocation"
)

func TestNewSQLStorage(t *testing.T) {
	connect := "root:hesoyam@tcp(127.0.0.1:3306)/dockertest"
	// Make sure DB and Table are created.

	storage, err := NewSQLStorage(connect)

	if err != nil {
		t.Fatal(err)
	}

	oloc := Geolocation{
		IP:        "127.42.24.1",
		CCode:     "IN",
		Country:   "India",
		City:      "Delhi",
		Longitude: -84.87503094689836,
		Latitude:  7.206435933364332,
		MValue:    7823011346,
	}

	err = storage.AddGeodata(oloc)

	if err != nil {
		t.Fatal(err)
	}

	gloc, err := storage.GetGeodata(oloc.IP)

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
