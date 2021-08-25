package models

import "time"

// import "errors"

type Geolocation struct {
	IP        string  `json:"ip"`
	CCode     string  `json:"country_code"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	MValue    int64   `json:"mystery_value"`
}

type Metrics struct {
	Duration time.Duration
	Total    int
	Rejected int
	Imported int
}

// Describes outcomes of parsing/importing a csv row.
type ImportEvent int

const (
	// Imported Successfully
	Success ImportEvent = iota
	AlreadyExists
	Invalid
	Failed
)

func (ie ImportEvent) GetStatus() string {
	switch ie {
	case Success:
		return "Successfully Imported"
	case AlreadyExists:
		return "Record Already Exists"
	case Invalid:
		return "Invalid Record Data"
	case Failed:
		return "Failed Importing Record"
	default:
		return "Unknown Event"
	}
}
