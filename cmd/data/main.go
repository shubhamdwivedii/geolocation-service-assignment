package main

import (
	"log"
	"os"

	importers "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/importers"
	sqlStorage "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/storage"
)

// This will populate DB with sample data
func main() {
	log.Println("Importing DB Data.....")
	DB_URL := os.Getenv("DB_URL")
	// DB_URL := "root:admin@tcp(127.0.0.1:3306)/geolocation"
	storage, _ := sqlStorage.NewSQLStorage(DB_URL)

	pwd, _ := os.Getwd()
	path := pwd + "/assignment/sample.csv"
	// Assuming this is run from project root folder.

	// Modify to take path as command-line argument

	csvImporter := importers.NewCSVImporter()

	importChannel, metricsChannel, err := csvImporter.Import(path)

	if err != nil {
		log.Fatal("Error parsing CSV", err.Error())
	}

	for glocation := range importChannel {
		log.Println("Importing::", glocation)
		err := storage.AddGeodata(glocation)
		if err != nil {
			log.Println("Error adding to storage:", err.Error())
		}
	}

	metrics := <-metricsChannel
	log.Println("METRICS", metrics)

	// function would not exit if run in go routine FIX LATER
	log.Println("End Of Sample Data Reached...")
}
