package service

import (
	"strings"

	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/geolocation"
	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/storage"
)

type Service interface {
	AddGeodata(Geolocation) error
	GetGeodata(string) (*Geolocation, error)
	GetAllByCCode(string) ([]*Geolocation, error)
}

type service struct {
	st Storage
}

func NewService(st Storage) Service {
	return &service{st}
}

func (s *service) AddGeodata(gloc Geolocation) error {
	return s.st.AddGeodata(gloc)
}

func (s *service) GetGeodata(ip string) (*Geolocation, error) {
	return s.st.GetGeodata(ip)
}

func (s *service) GetAllByCCode(ccode string) ([]*Geolocation, error) {
	locations, err := s.st.GetAllByCCode(ccode)

	if err != nil {
		return nil, err
	}

	// To emulate some business logic, converting City and Country to ALL UPPERCASE.
	for _, loc := range locations {
		(*loc).Country = strings.ToUpper((*loc).Country)
		(*loc).City = strings.ToUpper((*loc).City)
	}

	return locations, nil
}
