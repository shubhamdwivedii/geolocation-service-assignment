package geolocation

type Geolocation struct {
	IP        string  `json:"ip"`
	CCode     string  `json:"country_code"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	MValue    int64   `json:"mystery_value"`
}
