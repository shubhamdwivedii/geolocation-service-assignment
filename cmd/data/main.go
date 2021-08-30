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
	// DB_URL := os.Getenv("DB_URL")
	DB_URL := "root:hesoyam@tcp(127.0.0.1:3306)/dockertest"

	storage, err := sqlStorage.NewSQLStorage(DB_URL)

	if err != nil {
		log.Fatal("Error Initializing Storage...", err.Error())
	}

	pwd, _ := os.Getwd()
	// path := pwd + "/assignment/sample.csv"
	pathArg := os.Args[1:]
	if len(pathArg) <= 0 {
		log.Fatal("No CSV File Path Given To Import...")
	}
	path := pwd + pathArg[0]

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
