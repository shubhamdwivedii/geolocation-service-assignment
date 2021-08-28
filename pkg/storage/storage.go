package storage

import (
	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/geolocation"
)

type Storage interface {
	AddGeodata(Geolocation) error
	GetGeodata(string) (*Geolocation, error)
}
