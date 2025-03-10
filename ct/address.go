package ct

import (
	"github.com/mansio-gmbh/goapiutils/equals"
	"github.com/mansio-gmbh/goapiutils/hash"
)

type Address struct {
	Street                     *string        `json:"street,omitempty" dynamodbav:"street,omitempty" `
	HouseNumber                *string        `json:"houseNumber,omitempty" dynamodbav:"houseNumber,omitempty"`
	City                       *string        `json:"city,omitempty" dynamodbav:"city,omitempty" `
	PostalCode                 string         `json:"postalCode,omitempty" dynamodbav:"postalCode,omitempty"`
	CountryCode                string         `json:"countryCode,omitempty" dynamodbav:"countryCode,omitempty"`
	Name                       *string        `json:"name,omitempty" dynamodbav:"name,omitempty"`
	Name2                      *string        `json:"name2,omitempty" dynamodbav:"name2,omitempty"`
	Name3                      *string        `json:"name3,omitempty" dynamodbav:"name3,omitempty"`
	State                      *string        `json:"state,omitempty" dynamodbav:"state,omitempty"`
	TownArea                   *string        `json:"townArea,omitempty" dynamodbav:"townArea,omitempty"`
	EmailAddress               *string        `json:"emailAddress,omitempty" dynamodbav:"emailAddress,omitempty"`
	PhoneNumber                *string        `json:"phoneNumber,omitempty" dynamodbav:"phoneNumber,omitempty"`
	MobilePhoneNumber          *string        `json:"mobilePhoneNumber,omitempty" dynamodbav:"mobilePhoneNumber,omitempty"`
	AdditionalAddressLines     []string       `json:"additionalAddressLines,omitempty" dynamodbav:"additionalAddressLines,omitempty"`
	Reference                  *string        `json:"reference,omitempty" dynamodbav:"reference,omitempty"`
	ContactPerson              *string        `json:"contactPerson,omitempty" dynamodbav:"contactPerson,omitempty"`
	FaxNumber                  *string        `json:"faxNumber,omitempty" dynamodbav:"faxNumber,omitempty"`
	Gate                       *string        `json:"gate,omitempty" dynamodbav:"gate,omitempty"`
	Remarks                    *string        `json:"remarks,omitempty" dynamodbav:"remarks,omitempty"`
	LoadingWindows             LoadingWindows `json:"loadingWindows,omitempty" dynamodbav:"loadingWindows,omitempty"`
	Website                    *string        `json:"website,omitempty" dynamodbav:"website,omitempty"`
	VatID                      *string        `json:"vatID,omitempty" dynamodbav:"vatID,omitempty"`
	BusinessRegistrationNumber *string        `json:"businessRegistrationNumber,omitempty" dynamodbav:"businessRegistrationNumber,omitempty"`
	DistrictCourt              *string        `json:"districtCourt,omitempty" dynamodbav:"districtCourt,omitempty"`
	BuyerReference             *string        `json:"buyerReference,omitempty" dynamodbav:"buyerReference,omitempty"`
	EULicenseNumber            *string        `json:"euLicenseNumber,omitempty" dynamodbav:"euLicenseNumber,omitempty"`
}

func (a Address) IsSamePlace(other Address) bool {
	if a.CountryCode != other.CountryCode {
		return false
	}
	if a.PostalCode != other.PostalCode {
		return false
	}
	if !equals.Ptr(a.City, other.City) {
		return false
	}
	if !equals.Ptr(a.Street, other.Street) {
		return false
	}
	if !equals.Ptr(a.HouseNumber, other.HouseNumber) {
		return false
	}
	return true
}

func (a Address) IsEqual(other *Address) bool {
	if other == nil {
		return false
	}

	if !equals.Ptr(a.Street, other.Street) {
		return false
	}
	if !equals.Ptr(a.HouseNumber, other.HouseNumber) {
		return false
	}
	if !equals.Ptr(a.City, other.City) {
		return false
	}
	if a.PostalCode != other.PostalCode {
		return false
	}
	if a.CountryCode != other.CountryCode {
		return false
	}
	if !equals.Ptr(a.Name, other.Name) {
		return false
	}
	if !equals.Ptr(a.Name2, other.Name2) {
		return false
	}
	if !equals.Ptr(a.Name3, other.Name3) {
		return false
	}
	if !equals.Ptr(a.TownArea, other.TownArea) {
		return false
	}
	if !equals.Ptr(a.EmailAddress, other.EmailAddress) {
		return false
	}
	if !equals.Ptr(a.PhoneNumber, other.PhoneNumber) {
		return false
	}
	if !equals.Ptr(a.MobilePhoneNumber, other.MobilePhoneNumber) {
		return false
	}
	if !equals.Arr(a.AdditionalAddressLines, other.AdditionalAddressLines) {
		return false
	}
	if !equals.Ptr(a.Reference, other.Reference) {
		return false
	}
	if !equals.Ptr(a.ContactPerson, other.ContactPerson) {
		return false
	}
	if !equals.Ptr(a.FaxNumber, other.FaxNumber) {
		return false
	}
	if !equals.Ptr(a.Gate, other.Gate) {
		return false
	}
	if !equals.Ptr(a.Remarks, other.Remarks) {
		return false
	}
	if !equals.ArrEq(a.LoadingWindows, other.LoadingWindows) {
		return false
	}
	return true
}

func (a Address) ToLocation() *Location {
	return &Location{
		Address: &a,
	}
}

func (a Address) WithCoordinates(c Coordinates) *Location {
	return &Location{
		Address:     &a,
		Coordinates: &c,
	}
}

func (a Address) UniqueHash() (string, error) {
	return hash.SHA256(a)
}

func (a Address) MustUniqueHash() string {
	return hash.MustSHA256(a)
}

func (a Address) IsEmpty() bool {
	return (a.Street == nil || *a.Street == "") &&
		(a.HouseNumber == nil || *a.HouseNumber == "") &&
		(a.City == nil || *a.City == "") &&
		a.PostalCode == "" &&
		a.CountryCode == ""
}
