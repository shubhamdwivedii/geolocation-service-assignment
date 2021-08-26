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
	log.Println("Importing DB Data.....")
	DB_URL := os.Getenv("DB_URL")
	// DB_URL := "root:hesoyam@tcp(127.0.0.1:3306)/geolocation"
	storage, _ := sql.NewStorage(DB_URL)

	service := storageServices.NewService(storage)

	pwd, _ := os.Getwd()
	path := pwd + "/assignment/sample.csv"
	// Assuming this is run from project root folder.

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
