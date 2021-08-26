package main

import (
	"fmt"
	"log"
	"net/http"

	httpHandler "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/http"
	storageServices "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/services"
	"github.com/shubhamdwivedii/geolocation-service-assignment/pkg/storage/sql"
)

func main() {
	storage, _ := sql.NewStorage()

	service := storageServices.NewService(storage)

	handler := httpHandler.NewHandler(service)

	port := ":8080"

	http.Handle("/geodata/", handler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Geolocation Service!")
	})

	fmt.Println("Listening on Port", port)
	err := http.ListenAndServe(port, nil)
	log.Fatal(err)
}
