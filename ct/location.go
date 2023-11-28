package ct

type Location struct {
	Address     *Address     `json:"address" dynamodbav:"address" validate:"required"`
	Coordinates *Coordinates `json:"coordinates" dynamodbav:"coordinates" validate:"required"`
}
