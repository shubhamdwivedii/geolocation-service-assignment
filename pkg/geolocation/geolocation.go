package geolocation

import (
	"errors"
	"regexp"
)

type Geolocation struct {
	IP        string  `json:"ip"`
	CCode     string  `json:"country_code"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	MValue    int64   `json:"mystery_value"`
}

func validateIP(ip string) bool {
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	return re.MatchString(ip)
}

func validateCC(cc string) bool {
	re, _ := regexp.Compile(`^[A-Z][A-Z]$`)
	return re.MatchString(cc)
}

func ValidateGeolocation(gloc Geolocation) error {
	if !validateIP(gloc.IP) {
		return errors.New("Invalid IP Address.")
	}

	if !validateCC(gloc.CCode) {
		return errors.New("Invalid Country Code.")
	}

	if len(gloc.Country) > 20 || len(gloc.Country) < 2 {
		return errors.New("Invalid Country Name.")
	}

	if len(gloc.City) > 20 || len(gloc.City) < 2 {
		return errors.New("Invalid City Name.")
	}

	if gloc.Latitude > float64(90) || gloc.Latitude < float64(-90) {
		return errors.New("Invalid Latitude Value.")
	}

	if gloc.Longitude > float64(180) || gloc.Longitude < float64(-180) {
		return errors.New("Invalid Longitude Value.")
	}

	return nil
}
