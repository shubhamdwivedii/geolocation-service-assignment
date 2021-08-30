package geolocation_test

import (
	"testing"

	. "github.com/shubhamdwivedii/geolocation-service-assignment/pkg/geolocation"
)

func TestValidateGeolocation(t *testing.T) {
	tests := []struct {
		name string
		gloc Geolocation
		want bool
	}{
		{
			name: "Valid Geolocation",
			gloc: Geolocation{
				IP:        "200.106.141.15",
				CCode:     "SI",
				Country:   "Nepal",
				City:      "DuBuquemouth",
				Longitude: -84.87503094689836,
				Latitude:  7.206435933364332,
				MValue:    7823011346,
			},
			want: true,
		},
		{
			name: "Invalid Country Code",
			gloc: Geolocation{
				IP:        "160.103.7.140",
				CCode:     "CZA", // Should fail due to 3 letter CCode
				Country:   "Nicaragua",
				City:      "New Neva",
				Longitude: -68.31023296602508,
				Latitude:  -37.62435199624531,
				MValue:    7301823115,
			},
			want: false,
		},
		{
			name: "Invalid Longitude",
			gloc: Geolocation{
				IP:        "70.95.73.73",
				CCode:     "TL",
				Country:   "Saudi Arabia",
				City:      "Gradymouth",
				Longitude: -249.16675918861615, // Should fail due to invalid longitude value.
				Latitude:  -86.05920084416894,
				MValue:    2559997162,
			},
			want: false,
		},
		{
			name: "Invalid IP Address",
			gloc: Geolocation{
				IP:        "125.159.020.54", // Should fail due to invalid IP (020)
				CCode:     "LI",
				Country:   "Guyana",
				City:      "Port Karson",
				Longitude: -78.2274228596799,
				Latitude:  -163.26218895343357,
				MValue:    1337885276,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		// t.Run() runs a sub-test for each test case.
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGeolocation(tt.gloc)
			valid := true
			if err != nil {
				valid = false
			}

			if valid != tt.want {
				t.Errorf("Expected Geolocation Validity to be %v, But Got %v", tt.want, valid)
			}
		})
	}

}
