package types

type Unit struct {
	Unit  string  `json:"unit" mapstructure:"unit"`
	Value float64 `json:"value" mapstructure:"value"`
}

func (unit1 *Unit) Add(unit2 *Unit) *Unit {
	if unit2 == nil {
		return unit1
	}

	if unit1.Unit != unit2.Unit {
		return unit1
	}

	return &Unit{
		Value: unit1.Value + unit2.Value,
		Unit:  unit1.Unit,
	}
}

func AddUnit(unit1, unit2 *Unit) *Unit {
	if unit1 == nil {
		return unit2
	}

	if unit2 == nil {
		return unit1
	}

	return unit1.Add(unit2)
}
