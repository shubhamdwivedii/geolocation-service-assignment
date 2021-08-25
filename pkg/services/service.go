package service

import (
	"github.com/shubhamdwivedii/geolocation-service-assignment/models"
)

type Service interface {
	AddGeodata(models.Geolocation) error
	GetGeodata(string) (*models.Geolocation, error)
}

type StorageService interface {
	AddGeodata(models.Geolocation) error
	GetGeodata(string) (*models.Geolocation, error)
}

type service struct {
	store StorageService
}

func NewService(storage StorageService) Service {
	return &service{storage}
}

func (s *service) AddGeodata(gloc models.Geolocation) error {
	return s.store.AddGeodata(gloc)
}

func (s *service) GetGeodata(ip string) (*models.Geolocation, error) {
	return s.store.GetGeodata(ip)
}
