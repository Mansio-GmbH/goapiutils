package ct

type Address struct {
	Street                 *string  `json:"street,omitempty" dynamodbav:"street,omitempty" `
	HouseNumber            *string  `json:"houseNumber,omitempty" dynamodbav:"houseNumber,omitempty"`
	City                   *string  `json:"city,omitempty" dynamodbav:"city,omitempty" `
	PostalCode             string   `json:"postalCode,omitempty" dynamodbav:"postalCode,omitempty"`
	CountryCode            string   `json:"countryCode,omitempty" dynamodbav:"countryCode,omitempty"`
	Name                   *string  `json:"name,omitempty" dynamodbav:"name,omitempty"`
	Name2                  *string  `json:"name2,omitempty" dynamodbav:"name2,omitempty"`
	Name3                  *string  `json:"name3,omitempty" dynamodbav:"name3,omitempty"`
	TownArea               *string  `json:"townArea,omitempty" dynamodbav:"townArea,omitempty"`
	EmailAddress           *string  `json:"emailAddress,omitempty" dynamodbav:"emailAddress,omitempty"`
	PhoneNumber            *string  `json:"phoneNumber,omitempty" dynamodbav:"phoneNumber,omitempty"`
	MobilePhoneNumber      *string  `json:"mobilePhoneNumber,omitempty" dynamodbav:"mobilePhoneNumber,omitempty"`
	AdditionalAddressLines []string `json:"additionalAddressLines,omitempty" dynamodbav:"additionalAddressLines,omitempty"`
	Reference              *string  `json:"reference,omitempty" dynamodbav:"reference,omitempty"`
	ContactPerson          *string  `json:"contactPerson,omitempty" dynamodbav:"contactPerson,omitempty"`
	FaxNumber              *string  `json:"faxNumber,omitempty" dynamodbav:"faxNumber,omitempty"`
	Gate                   *string  `json:"gate,omitempty" dynamodbav:"gate,omitempty"`
	Remarks                *string  `json:"remarks,omitempty" dynamodbav:"remarks,omitempty"`
}
