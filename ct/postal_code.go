package ct

type PostalCode struct {
	PostalCode  string `json:"postalCode" dynamodbav:"postalCode"`
	CountryCode string `json:"countryCode" dynamodbav:"countryCode"`
}

func (pc PostalCode) String() string {
	return pc.CountryCode + pc.PostalCode
}
