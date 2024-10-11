package ct_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/mansio-gmbh/goapiutils/ct"
	"github.com/stretchr/testify/require"
)

func TestLocationSamePlace(t *testing.T) {
	tests := []struct {
		location1 ct.Location
		location2 ct.Location
		samePlace bool
	}{
		{
			location1: ct.Location{
				Coordinates: &ct.Coordinates{Latitude: 50.76770455131623, Longitude: 6.098805099158805},
			},
			location2: ct.Location{
				Coordinates: &ct.Coordinates{Latitude: 50.76770455131623, Longitude: 6.098805099158805},
			},
			samePlace: true,
		},
		{
			location1: ct.Location{
				Address: &ct.Address{
					Street:     aws.String("Musterstr 2"),
					PostalCode: "12345",
				},
			},
			location2: ct.Location{
				Address: &ct.Address{
					Street:     aws.String("Musterstr 2"),
					PostalCode: "12345",
				},
			},
			samePlace: true,
		},
		{
			location1: ct.Location{
				Address: &ct.Address{
					Street:      aws.String("Musterstr 2"),
					PostalCode:  "12345",
					CountryCode: "DE",
				},
			},
			location2: ct.Location{
				Address: &ct.Address{
					Street:      aws.String("Musterstr 2"),
					PostalCode:  "12345",
					CountryCode: "EN",
				},
			},
			samePlace: false,
		},
	}

	for _, test := range tests {
		require.Equal(t, test.samePlace, test.location1.IsSamePlace(test.location2), "Expected to be same place: %v and %v", test.location1, test.location2)
	}
}
