package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	st "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/storage"
)

type HttpHandler struct {
	storage st.Storage
}

// HttpHandler implements http.Handler Interface
/*
	http.Handler interafce {
		ServeHTTP(http.RespohnseWriter, *http.Request)
	}
*/

func NewHandler(storage st.Storage) http.Handler {
	var handler = HttpHandler{
		storage: storage,
	}
	return &handler
}

func (handler *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handler.get(w, r)
	default:
		respondWithJSON(w, http.StatusMethodNotAllowed, "Invalid Method.")
	}
}

func (handler *HttpHandler) get(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.String(), "/")
	// SSSSSS r.URL.Query()
	ip := segs[len(segs)-1]

	geolocation, err := handler.storage.GetGeodata(ip)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusFound, geolocation)
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
