package parser

import (
	"errors"

	"github.com/shubhamdwivedii/geolocation-service-assignment/models"
	"github.com/shubhamdwivedii/geolocation-service-assignment/pkg/parsers/csv"
)

type Parser interface {
	Import(string) (<-chan models.Geolocation, error)
	GetMetrics() models.Metrics
}

func NewParser(ptype string) (Parser, error) {
	if ptype == "csv" {
		var parser *csv.CSVParser
		parser = new(csv.CSVParser)
		return parser, nil
	}
	return nil, errors.New("Invalid Parser Type")
}
