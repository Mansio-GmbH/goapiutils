package ct

import "github.com/mansio-gmbh/goapiutils/equals"

type UnitValue struct {
	Unit  *string `json:"unit" dynamodbav:"unit,omitempty"`
	Value float64 `json:"value" dynamodbav:"value"`
}

func NewUnitValue(value float64, unit string) *UnitValue {
	return &UnitValue{
		Unit:  &unit,
		Value: value,
	}
}

func (u UnitValue) IsEqual(other *UnitValue) bool {
	if other == nil {
		return false
	}
	return equals.Ptr(u.Unit, other.Unit) && u.Value == other.Value
}
