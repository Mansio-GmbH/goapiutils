package ct

import (
	"math"

	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/mansio-gmbh/goapiutils/hash"
)

type Coordinates struct {
	Latitude  float64 `json:"latitude" dynamodbav:"latitude" validate:"gte=-90,lte=90"`
	Longitude float64 `json:"longitude" dynamodbav:"longitude" validate:"gte=-180,lte=180"`
}

type CoordinatesWithTimestamp struct {
	Coordinates
	Timestamp *chrono.Time `json:"timestamp" dynamodbav:"timestamp"`
}

func (c Coordinates) IsZero() bool {
	return c.Latitude == 0 && c.Longitude == 0
}

func (c Coordinates) HaversineDistance(other Coordinates) float64 {
	const earthRadius = 6371.0
	lat1 := c.Latitude
	lon1 := c.Longitude
	lat2 := other.Latitude
	lon2 := other.Longitude

	lat1Rad := lat1 * (math.Pi / 180)
	lon1Rad := lon1 * (math.Pi / 180)
	lat2Rad := lat2 * (math.Pi / 180)
	lon2Rad := lon2 * (math.Pi / 180)

	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	a := (math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2))

	dc := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * dc
}

func (c Coordinates) HaversineDistanceAsKm(other Coordinates) *UnitValue {
	distance := c.HaversineDistance(other)
	return NewUnitValue(distance, "KM")
}

func (c Coordinates) HaversineDistanceAsMeter(other Coordinates) *UnitValue {
	distance := c.HaversineDistance(other) * 1000
	return NewUnitValue(distance, "M")
}

func (c Coordinates) IsNear(other Coordinates, maxDistance float64) bool {
	return c.HaversineDistance(other) <= maxDistance
}

func (c Coordinates) EqualByDistance(other Coordinates) bool {
	return c.IsNear(other, 0.001) // 1 meter
}

func (c Coordinates) IsEqual(other *Coordinates) bool {
	if other == nil {
		return false
	}
	return c.Latitude == other.Latitude && c.Longitude == other.Longitude
}

func (c Coordinates) ToLocation() *Location {
	return &Location{
		Coordinates: &c,
	}
}

func (c Coordinates) WithAddress(a Address) *Location {
	return &Location{
		Coordinates: &c,
		Address:     &a,
	}
}

func (c Coordinates) UniqueHash() (string, error) {
	return hash.SHA256(c)
}

func (c Coordinates) MustUniqueHash() string {
	return hash.MustSHA256(c)
}
