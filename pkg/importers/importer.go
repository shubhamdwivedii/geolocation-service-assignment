package importers

import (
	"time"

	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/geolocation"
)

type Metrics struct {
	Duration time.Duration
	Total    int
	Rejected int
	Imported int
}

type Importer interface {
	Import(string) (<-chan Geolocation, <-chan Metrics, error)
}
