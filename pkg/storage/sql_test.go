package storage

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/geolocation"
)

func TestNewSQLStorage(t *testing.T) {
	connect := "root:admin@tcp(127.0.0.1:3306)/dockertest"
	// Make sure DB and Table are created.

	storage, err := NewSQLStorage(connect)
	require.NoError(t, err)

	oloc := Geolocation{
		IP:        "127.42.24.1",
		CCode:     "IN",
		Country:   "India",
		City:      "Delhi",
		Longitude: -84.87503094689836,
		Latitude:  7.206435933364332,
		MValue:    7823011346,
	}

	err = storage.AddGeodata(oloc)

	if err != nil && strings.Contains(err.Error(), "1062") {
		log.Println("Duplicate Entry Error... Continuing With Test...")
	} else {
		require.NoError(t, err)
	}

	gloc, err := storage.GetGeodata(oloc.IP)
	require.NoError(t, err)

	assert.Equal(t, *gloc, oloc, "Expected Both To Be Same.")

	glocs, err := storage.GetAllByCCode("IN")
	require.NoError(t, err)

	for _, loc := range glocs {
		assert.Equal(t, (*loc).CCode, "IN", "Country Code Should Match.")
	}
}
