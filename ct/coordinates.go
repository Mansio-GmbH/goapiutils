package ct

type Coordinates struct {
	Latitude  float64 `json:"latitude" dynamodbav:"latitude" validate:"gte=-90,lte=90"`
	Longitude float64 `json:"longitude" dynamodbav:"longitude" validate:"gte=-180,lte=180"`
}
