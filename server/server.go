package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/shubhamdwivedii/geolocation-service-assignment/geolocation"
)

type GeoDataHandler struct {
	geodata *geolocation.GeoData
}

// Handler interface has a ServeHTTP method
func (gdh *GeoDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		gdh.get(w, r)
	// case "POST":
	// gdh.post(w, r)
	default:
		respondWithJSON(w, http.StatusMethodNotAllowed, "invalid method")
	}
}

func (gdh *GeoDataHandler) get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REQUEST", r.URL, r.URL.String())
	fmt.Println("REQ BODY", r.Body)
	segs := strings.Split(r.URL.String(), "/")
	ip := segs[len(segs)-1]

	geolocation, err := gdh.geodata.GetGeoData(ip)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, geolocation)
}

func respondWithJSON(w http.ResponseWriter, code int, data interface{}) {
	response, err := json.Marshal(data) // returns []byte (json in string)
	if err != nil {
		fmt.Println("Error Marshalling Response Data:", err.Error())
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	// As soon as w.Write() is executed, the Server will send the response.
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func RunServer(geodata *geolocation.GeoData) {
	port := ":8080"
	geoHandler := GeoDataHandler{geodata: geodata}
	// http.Handle("/geodata", &geoHandler)
	http.Handle("/geodata/", &geoHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Geolocation Service!")
	})

	fmt.Println("Listening on Port", port)
	err := http.ListenAndServe(port, nil)
	log.Fatal(err)
}
