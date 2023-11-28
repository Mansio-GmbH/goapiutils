package ct

type UnitValue struct {
	Unit  *string `json:"unit" dynamodbav:"unit,omitempty"`
	Value float64 `json:"value" dynamodbav:"value"`
}
