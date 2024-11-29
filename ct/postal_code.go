package ct

type PostalCode struct {
	PostalCode  string `json:"postalCode" dynamodbav:"postalCode"`
	CountryCode string `json:"countryCode" dynamodbav:"countryCode"`
}

func (pc PostalCode) String() string {
	return pc.CountryCode + pc.PostalCode
}

func (pc PostalCode) IsEqual(other *PostalCode) bool {
	if other == nil {
		return false
	}
	return pc.PostalCode == other.PostalCode && pc.CountryCode == other.CountryCode
}
