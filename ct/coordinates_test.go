package ct_test

import (
	"testing"

	"github.com/mansio-gmbh/goapiutils/ct"
	"github.com/stretchr/testify/require"
)

func TestCoordinates(t *testing.T) {
	tests := []struct {
		coord1          ct.Coordinates
		coord2          ct.Coordinates
		equal           bool
		equalByDistance bool
		distance        float64
		isNear          []ct.Coordinates
		isNotNear       []ct.Coordinates
	}{
		{
			coord1:          ct.Coordinates{Latitude: 50.76770455131623, Longitude: 6.098805099158805},
			coord2:          ct.Coordinates{Latitude: 50.76770455131623, Longitude: 6.098805099158805},
			equal:           true,
			equalByDistance: true,
			distance:        0,
			isNear:          []ct.Coordinates{{Latitude: 50.76770455131623, Longitude: 6.098805099158805}},
			isNotNear:       []ct.Coordinates{{Latitude: 50.86770455131623, Longitude: 6.098805099158805}},
		},
		{
			coord1:          ct.Coordinates{Latitude: 50.76770455131623, Longitude: 6.098805099158805},
			coord2:          ct.Coordinates{Latitude: 50.7677045426432, Longitude: 6.0988050993896},
			equal:           false,
			equalByDistance: true,
			distance:        0.00009,
		},
		{
			coord1:          ct.Coordinates{Latitude: 50.76770455131623, Longitude: 6.098805099158805},
			coord2:          ct.Coordinates{Latitude: 50.7676850426432, Longitude: 6.0986924463896},
			equal:           false,
			equalByDistance: false,
			distance:        0.00821,
		},
	}

	for idx, test := range tests {
		require.Equal(t, test.equal, test.coord1.Equal(test.coord2), "Test %d: Expected to be equal: %v and %v", idx, test.coord1, test.coord2)
		require.Equal(t, test.equalByDistance, test.coord1.EqualByDistance(test.coord2), "Test %d: Expected to be equal by distance: %v and %v", idx, test.coord1, test.coord2)
		require.InDelta(t, test.distance, test.coord1.HaversineDistance(test.coord2), 0.0001, "Test %d: Expected distance: %v and %v to be %f", idx, test.coord1, test.coord2, test.distance)
		for _, near := range test.isNear {
			require.True(t, test.coord1.IsNear(near, 10), "Test %d: Expected to be near: %v and %v", idx, test.coord1, near)
		}
		for _, notNear := range test.isNotNear {
			require.False(t, test.coord1.IsNear(notNear, 10), "Test %d: Expected to be not near: %v and %v", idx, test.coord1, notNear)
		}
	}
}
