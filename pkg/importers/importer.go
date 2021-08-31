package importers

import (
	"time"

	sv "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/service"
)

type Metrics struct {
	Duration time.Duration
	Total    int
	Rejected int
	Imported int
}

type Importer interface {
	Import(string, sv.Service) (<-chan Metrics, error)
}
