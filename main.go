package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shubhamdwivedii/geolocation-service-assignment/geolocation"
	"github.com/shubhamdwivedii/geolocation-service-assignment/server"
)

func main() {
	DB_URL := os.Getenv("DB_URL")
	// DB_URL := "root:admin@tcp(127.0.0.1:3306)/geolocation"

	geodata, err := geolocation.InitGeoData(DB_URL)
	if err != nil {
		log.Fatal("Error: Initializing GeoData", err.Error())
	}

	fmt.Println(geodata)

	pwd, _ := os.Getwd()
	path := pwd + "/assignment/sample.csv"
	// path = pwd + "/assignment/data_dump.csv"

	// url := "https://drive.google.com/uc?id=1G6ALooi2ba0WSQpv_3Q-lhInWJ7APEz6&export=download"
	metrics, err := geolocation.ReadCSVFile(path, geodata)

	fmt.Printf(`Imported CSV Data Successfully:
		Time: %v ms
		Total Records: %v
		Rejeced: %v
		Imported: %v`, metrics.Duration, metrics.Total, metrics.Rejected, metrics.Imported)

	fmt.Println("")
	server.RunServer(geodata)
}

// UNIT TEST
// SOLID KA S (SRP)
// squirrel query builder
// sql injection
