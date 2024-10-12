package ct

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

func (l Location) SetAddress(address *Address) *Location {
	l.Address = address
	return &l
}

func (l Location) SetCoordinates(coordinates *Coordinates) *Location {
	l.Coordinates = coordinates
	return &l
}
