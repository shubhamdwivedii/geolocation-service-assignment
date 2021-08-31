package main

import (
	"log"
	"os"

	importers "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/importers"
	sv "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/service"
	sqlStorage "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/storage"
)

// This will populate DB with sample data
func main() {
	log.Println("Importing DB Data.....")
	DB_URL := os.Getenv("DB_URL")
	// DB_URL := "root:admin@tcp(127.0.0.1:3306)/dockertest"

	storage, err := sqlStorage.NewSQLStorage(DB_URL)
	if err != nil {
		log.Fatal("Error Initializing Storage...", err.Error())
	}

	service := sv.NewService(storage)

	pwd, _ := os.Getwd()
	// path := pwd + "/assignment/sample.csv"
	pathArg := os.Args[1:]
	if len(pathArg) <= 0 {
		log.Fatal("No CSV File Path Given To Import...")
	}
	path := pwd + pathArg[0]

	csvImporter := importers.NewCSVImporter()
	metricsChannel, err := csvImporter.Import(path, service)

	if err != nil {
		log.Fatal("Error parsing CSV", err.Error())
	}

	metrics := <-metricsChannel
	log.Printf(`CSV Import Results:
	Duration: %v 
	Total Records: %v
	Rejected: %v
	Imported: %v`, metrics.Duration, metrics.Total, metrics.Rejected, metrics.Imported)
}
