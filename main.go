package main

import (
	"fmt"
	"log"

	"github.com/shubhamdwivedii/geolocation-service-assignment/geolocation"
	"github.com/shubhamdwivedii/geolocation-service-assignment/server"
)

func main() {
	fmt.Println("Hello Geolocation")
	geodata, err := geolocation.InitGeoData()
	server.RunServer(geodata)

	if err != nil {
		log.Fatal("Error: Initializing GeoData", err.Error())
	}
	geolocation.ReadCSV(geodata)
	g, err := geodata.GetGeoData("70.95.73.73")
	if err != nil {
		fmt.Println("Error getting geodata", err.Error())
	}
	fmt.Println(g)

}
