package ct

import "github.com/mansio-gmbh/goapiutils/hash"

type Location struct {
	Address     *Address     `json:"address" dynamodbav:"address"`
	Coordinates *Coordinates `json:"coordinates" dynamodbav:"coordinates"`
}

func (l Location) IsSamePlace(other Location) bool {
	if l.Address == nil && other.Address == nil && l.Coordinates == nil && other.Coordinates == nil {
		return true
	}
	if l.Address != nil && other.Address != nil {
		return l.Address.IsSamePlace(*other.Address)
	}
	if l.Coordinates != nil && other.Coordinates != nil {
		return l.Coordinates.EqualByDistance(*other.Coordinates)
	}
	return false
}

func (l Location) GetAddress() *Address {
	return l.Address
}

func (l Location) GetCoordinates() *Coordinates {
	return l.Coordinates
}

func (l *Location) SetAddress(address *Address) {
	l.Address = address
}

func (l *Location) SetCoordinates(coordinates *Coordinates) {
	l.Coordinates = coordinates
}

func (l Location) IsEqual(other *Location) bool {
	if other == nil {
		return false
	}
	if l.Address != nil && !l.Address.IsEqual(other.Address) {
		return false
	}

	if l.Coordinates != nil && !l.Coordinates.IsEqual(other.Coordinates) {
		return false
	}
	return true
}

func (l Location) UniqueHash() (string, error) {
	return hash.SHA256(l)
}

func (l Location) MustUniqueHash() string {
	return hash.MustSHA256(l)
}

func (l Location) IsEmpty() bool {
	if l.Address == nil && l.Coordinates == nil {
		return true
	}
	if l.Address != nil && !l.Address.IsEmpty() {
		return false
	}
	if l.Coordinates != nil && !l.Coordinates.IsEmpty() {
		return false
	}
	return true
}

func WrapAddress(addr Address) Location {
	return Location{
		Address: &addr,
	}
}

func WrapAddressPtr(addr *Address) Location {
	if addr == nil {
		return Location{}
	}
	return Location{
		Address: addr,
	}
}

func WrapAddresses(addr ...Address) []Location {
	locations := make([]Location, len(addr))
	for i := range addr {
		locations[i] = WrapAddress(addr[i])
	}
	return locations
}

func WrapCoordinate(coord Coordinates) Location {
	return Location{
		Coordinates: &coord,
	}
}

func WrapCoordinatePtr(coord *Coordinates) Location {
	if coord == nil {
		return Location{}
	}
	return Location{
		Coordinates: coord,
	}
}

func WrapCoordinates(coords ...Coordinates) []Location {
	locations := make([]Location, len(coords))
	for i := range coords {
		locations[i] = WrapCoordinate(coords[i])
	}
	return locations
}
