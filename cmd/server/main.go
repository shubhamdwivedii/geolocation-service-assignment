package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	httpHandler "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/http"
	sqlStorage "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/storage"
)

// This will run the http server.
func main() {
	log.Println("Starting Server.....")
	DB_URL := os.Getenv("DB_URL")
	// DB_URL := "root:admin@tcp(127.0.0.1:3306)/geolocation"

	// FIX THESE KIND OF ERROR IN RESPONSE.
	// {"error":"Error 1045: Access denied for user ''@'localhost' (using password: NO)"}

	storage, _ := sqlStorage.NewSQLStorage(DB_URL)
	handler := httpHandler.NewHandler(storage)

	port := ":8080"

	http.Handle("/geodata/", handler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Geolocation Service!")
	})

	fmt.Println("Listening on Port", port)
	err := http.ListenAndServe(port, nil)
	log.Fatal(err)
}
