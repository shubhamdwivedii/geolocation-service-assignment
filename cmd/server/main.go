package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	httpHandler "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/http"
	storageServices "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/services"
	"github.com/shubhamdwivedii/geolocation-service-assignment/pkg/storage/sql"
)

// This will run the http server.
func main() {
	log.Println("Starting Server.....")
	DB_URL := os.Getenv("DB_URL")
	// DB_URL := "root:hesoyam@tcp(127.0.0.1:3306)/geolocation"

	storage, _ := sql.NewStorage(DB_URL)

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
