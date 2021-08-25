package main

import (
	"log"
	"os"

	parser "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/parsers"
	storageServices "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/services"
	"github.com/shubhamdwivedii/geolocation-service-assignment/pkg/storage/sql"
)

// This will populate DB with sample data

func main() {
	storage, _ := sql.NewStorage()

	service := storageServices.NewService(storage)

	pwd, _ := os.Getwd()
	path := pwd[:len(pwd)-9] + "/assignment/sample.csv"
	// Fix path later.

	csvparser, err := parser.NewParser("csv")
	if err != nil {
		log.Fatal(err.Error())
	}

	importChannel, err := csvparser.Import(path)

	if err != nil {
		log.Fatal("Error parsing CSV", err.Error())
	}

	// go func() {
	for glocation := range importChannel {
		log.Println("Importing::", glocation)
		err := service.AddGeodata(glocation)
		if err != nil {
			log.Println("Error adding to storage:", err.Error())
		}
	}
	// }()

	// function would exit if run in go routine FIX LATER

	log.Println("End Of Sample Data Reached...")
}
