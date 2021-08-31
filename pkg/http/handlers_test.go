package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/geolocation"
	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	storage, _ := NewMockStorage()

	rr := httptest.NewRecorder()

	handler := NewHandler(storage)

	oloc := Geolocation{
		IP:        "127.42.24.1",
		CCode:     "IN",
		Country:   "India",
		City:      "Delhi",
		Longitude: -84.87503094689836,
		Latitude:  7.206435933364332,
		MValue:    7823011346,
	}

	storage.AddGeodata(oloc)

	r, err := http.NewRequest(http.MethodGet, "/geodata/127.42.24.1", nil)
	require.NoError(t, err)

	handler.ServeHTTP(rr, r)

	rs := rr.Result()

	body, err := ioutil.ReadAll(rs.Body)
	require.NoError(t, err)

	var gloc Geolocation
	err = json.Unmarshal(body, &gloc)
	require.NoError(t, err)

	assert.Equal(t, gloc, oloc, "They Should Be Equal.")
}

/*********** MOCK DB Storage ***********/
type MockStorage struct {
	db map[string]Geolocation
}

func NewMockStorage() (Storage, error) {
	s := new(MockStorage)
	s.db = make(map[string]Geolocation)
	return s, nil
}

func (s *MockStorage) AddGeodata(gloc Geolocation) error {
	err := ValidateGeolocation(gloc)
	if err != nil {
		return err
	}

	if _, ok := s.db[gloc.IP]; ok {
		return errors.New("Duplicate Entry.")
	} else {
		s.db[gloc.IP] = gloc
	}
	return nil
}

func (s *MockStorage) GetGeodata(ip string) (*Geolocation, error) {
	if gloc, ok := s.db[ip]; !ok {
		return nil, errors.New("Does Not Exits In DB.")
	} else {
		return &gloc, nil
	}
}
