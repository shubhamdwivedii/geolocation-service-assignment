package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shubhamdwivedii/geolocation-service-assignment/geolocation"
	"github.com/shubhamdwivedii/geolocation-service-assignment/server"
)

func main() {
	fmt.Println("Hello Geolocation")
	geodata, err := geolocation.InitGeoData()
	if err != nil {
		log.Fatal("Error: Initializing GeoData", err.Error())
	}

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
