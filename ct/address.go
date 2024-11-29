package ct

type Address struct {
	Street                 *string        `json:"street,omitempty" dynamodbav:"street,omitempty" `
	HouseNumber            *string        `json:"houseNumber,omitempty" dynamodbav:"houseNumber,omitempty"`
	City                   *string        `json:"city,omitempty" dynamodbav:"city,omitempty" `
	PostalCode             string         `json:"postalCode,omitempty" dynamodbav:"postalCode,omitempty"`
	CountryCode            string         `json:"countryCode,omitempty" dynamodbav:"countryCode,omitempty"`
	Name                   *string        `json:"name,omitempty" dynamodbav:"name,omitempty"`
	Name2                  *string        `json:"name2,omitempty" dynamodbav:"name2,omitempty"`
	Name3                  *string        `json:"name3,omitempty" dynamodbav:"name3,omitempty"`
	TownArea               *string        `json:"townArea,omitempty" dynamodbav:"townArea,omitempty"`
	EmailAddress           *string        `json:"emailAddress,omitempty" dynamodbav:"emailAddress,omitempty"`
	PhoneNumber            *string        `json:"phoneNumber,omitempty" dynamodbav:"phoneNumber,omitempty"`
	MobilePhoneNumber      *string        `json:"mobilePhoneNumber,omitempty" dynamodbav:"mobilePhoneNumber,omitempty"`
	AdditionalAddressLines []string       `json:"additionalAddressLines,omitempty" dynamodbav:"additionalAddressLines,omitempty"`
	Reference              *string        `json:"reference,omitempty" dynamodbav:"reference,omitempty"`
	ContactPerson          *string        `json:"contactPerson,omitempty" dynamodbav:"contactPerson,omitempty"`
	FaxNumber              *string        `json:"faxNumber,omitempty" dynamodbav:"faxNumber,omitempty"`
	Gate                   *string        `json:"gate,omitempty" dynamodbav:"gate,omitempty"`
	Remarks                *string        `json:"remarks,omitempty" dynamodbav:"remarks,omitempty"`
	LoadingWindows         LoadingWindows `json:"loadingWindows,omitempty" dynamodbav:"loadingWindows,omitempty"`
}

func (a Address) IsSamePlace(other Address) bool {
	if a.CountryCode != other.CountryCode {
		return false
	}
	if a.PostalCode != other.PostalCode {
		return false
	}
	if !ptrEq(a.City, other.City) {
		return false
	}
	if !ptrEq(a.Street, other.Street) {
		return false
	}
	if !ptrEq(a.HouseNumber, other.HouseNumber) {
		return false
	}
	return true
}

func ptrEq[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func stringSliceEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (a Address) IsEqual(other *Address) bool {
	if other == nil {
		return false
	}

	if !ptrEq(a.Street, other.Street) {
		return false
	}
	if !ptrEq(a.HouseNumber, other.HouseNumber) {
		return false
	}
	if !ptrEq(a.City, other.City) {
		return false
	}
	if a.PostalCode != other.PostalCode {
		return false
	}
	if a.CountryCode != other.CountryCode {
		return false
	}
	if !ptrEq(a.Name, other.Name) {
		return false
	}
	if !ptrEq(a.Name2, other.Name2) {
		return false
	}
	if !ptrEq(a.Name3, other.Name3) {
		return false
	}
	if !ptrEq(a.TownArea, other.TownArea) {
		return false
	}
	if !ptrEq(a.EmailAddress, other.EmailAddress) {
		return false
	}
	if !ptrEq(a.PhoneNumber, other.PhoneNumber) {
		return false
	}
	if !ptrEq(a.MobilePhoneNumber, other.MobilePhoneNumber) {
		return false
	}
	if !stringSliceEq(a.AdditionalAddressLines, other.AdditionalAddressLines) {
		return false
	}
	if !ptrEq(a.Reference, other.Reference) {
		return false
	}
	if !ptrEq(a.ContactPerson, other.ContactPerson) {
		return false
	}
	if !ptrEq(a.FaxNumber, other.FaxNumber) {
		return false
	}
	if !ptrEq(a.Gate, other.Gate) {
		return false
	}
	if !ptrEq(a.Remarks, other.Remarks) {
		return false
	}
	if !a.LoadingWindows.IsEqual(other.LoadingWindows) {
		return false
	}
	return true
}
