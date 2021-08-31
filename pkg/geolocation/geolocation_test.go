package geolocation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateGeolocation(t *testing.T) {
	tests := []struct {
		name string
		gloc Geolocation
		err  bool
	}{
		{
			name: "Valid Geolocation",
			gloc: Geolocation{
				IP:        "200.106.141.15",
				CCode:     "SI",
				Country:   "Nepal",
				City:      "DuBuquemouth",
				Latitude:  7.206435933364332,
				Longitude: -84.87503094689836,
				MValue:    7823011346,
			},
			err: false,
		},
		{
			name: "Invalid Country Code",
			gloc: Geolocation{
				IP:        "160.103.7.140",
				CCode:     "CZA", // Should fail due to 3 letter CCode
				Country:   "Nicaragua",
				City:      "New Neva",
				Latitude:  -37.62435199624531,
				Longitude: -68.31023296602508,
				MValue:    7301823115,
			},
			err: true,
		},
		{
			name: "Invalid Longitude",
			gloc: Geolocation{
				IP:        "70.95.73.73",
				CCode:     "TL",
				Country:   "Saudi Arabia",
				City:      "Gradymouth",
				Latitude:  -86.05920084416894,
				Longitude: -249.16675918861615, // Should fail due to invalid longitude value.
				MValue:    2559997162,
			},
			err: true,
		},
		{
			name: "Invalid IP Address",
			gloc: Geolocation{
				IP:        "125.159.020.54", // Should fail due to invalid IP (020)
				CCode:     "LI",
				Country:   "Guyana",
				City:      "Port Karson",
				Latitude:  -163.26218895343357,
				Longitude: -78.2274228596799,
				MValue:    1337885276,
			},
			err: true,
		},
	}

	for _, tt := range tests {
		// t.Run() runs a sub-test for each test case.
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGeolocation(tt.gloc)
			if tt.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
