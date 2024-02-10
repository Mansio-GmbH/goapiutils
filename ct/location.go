package ct

type Location struct {
	Address     *Address     `json:"address" dynamodbav:"address"`
	Coordinates *Coordinates `json:"coordinates" dynamodbav:"coordinates"`
}
