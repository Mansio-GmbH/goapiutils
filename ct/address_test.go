package ct_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/mansio-gmbh/goapiutils/ct"
	"github.com/stretchr/testify/require"
)

func TestAddressIsSamePlace(t *testing.T) {
	tests := []struct {
		address1  ct.Address
		address2  ct.Address
		samePlace bool
	}{
		{
			address1: ct.Address{
				Street: aws.String("Musterstr 2"),
			},
			address2: ct.Address{
				Street: aws.String("Musterstr 2"),
			},
			samePlace: true,
		},
		{
			address1: ct.Address{
				Street:     aws.String("Musterstr 2"),
				PostalCode: "12345",
			},
			address2: ct.Address{
				Street:     aws.String("Musterstr 2"),
				PostalCode: "12345",
			},
			samePlace: true,
		},
	}
	for idx, test := range tests {
		require.Equal(t, test.samePlace, test.address1.IsSamePlace(test.address2), "Test %d: Expected to be same place: %v and %v", idx, test.address1, test.address2)
	}
}

func TestAddressIsEqual(t *testing.T) {
	tests := []struct {
		address1 ct.Address
		address2 ct.Address
		equal    bool
	}{
		{
			address1: ct.Address{
				Street:            aws.String("Musterstr 2"),
				HouseNumber:       aws.String("2"),
				City:              aws.String("Musterstadt"),
				PostalCode:        "12345",
				CountryCode:       "DE",
				Name:              aws.String("Name"),
				Name2:             aws.String("Name2"),
				Name3:             aws.String("Name3"),
				TownArea:          aws.String("TownArea"),
				EmailAddress:      aws.String("foobar@baz.com"),
				PhoneNumber:       aws.String("1234567890"),
				MobilePhoneNumber: aws.String("0987654321"),
				AdditionalAddressLines: []string{
					"AdditionalAddressLine1",
					"AdditionalAddressLine2",
				},
				Reference:     aws.String("Reference"),
				ContactPerson: aws.String("ContactPerson"),
				FaxNumber:     aws.String("0987654321"),
				Gate:          aws.String("Gate"),
				Remarks:       aws.String("Remarks"),
				LoadingWindows: ct.LoadingWindows{
					{
						StartsAt: "08:00",
						EndsAt:   "12:00",
					},
				},
			},
			address2: ct.Address{
				Street:            aws.String("Musterstr 2"),
				HouseNumber:       aws.String("2"),
				City:              aws.String("Musterstadt"),
				PostalCode:        "12345",
				CountryCode:       "DE",
				Name:              aws.String("Name"),
				Name2:             aws.String("Name2"),
				Name3:             aws.String("Name3"),
				TownArea:          aws.String("TownArea"),
				EmailAddress:      aws.String("foobar@baz.com"),
				PhoneNumber:       aws.String("1234567890"),
				MobilePhoneNumber: aws.String("0987654321"),
				AdditionalAddressLines: []string{
					"AdditionalAddressLine1",
					"AdditionalAddressLine2",
				},
				Reference:     aws.String("Reference"),
				ContactPerson: aws.String("ContactPerson"),
				FaxNumber:     aws.String("0987654321"),
				Gate:          aws.String("Gate"),
				Remarks:       aws.String("Remarks"),
				LoadingWindows: ct.LoadingWindows{
					{
						StartsAt: "08:00",
						EndsAt:   "12:00",
					},
				},
			},
			equal: true,
		},
		{
			address1: ct.Address{
				Street:            aws.String("Musterstr 2"),
				HouseNumber:       aws.String("2"),
				City:              aws.String("Musterstadt"),
				PostalCode:        "12345",
				CountryCode:       "DE",
				Name:              aws.String("Name"),
				Name2:             aws.String("Name2"),
				Name3:             aws.String("Name3"),
				TownArea:          aws.String("TownArea"),
				EmailAddress:      aws.String("foobar@baz.com"),
				PhoneNumber:       aws.String("1234567890"),
				MobilePhoneNumber: aws.String("0987654321"),
				AdditionalAddressLines: []string{
					"AdditionalAddressLine1",
					"AdditionalAddressLine2",
				},
				Reference:     aws.String("Reference"),
				ContactPerson: aws.String("ContactPerson"),
				FaxNumber:     aws.String("0987654321"),
				Gate:          aws.String("Gate"),
				Remarks:       aws.String("Remarks"),
				LoadingWindows: ct.LoadingWindows{
					{
						StartsAt: "08:00",
						EndsAt:   "12:00",
					},
				},
			},
			address2: ct.Address{
				Street:            aws.String("Musterstr 2"),
				HouseNumber:       aws.String("2"),
				City:              aws.String("Musterstadt"),
				PostalCode:        "12345",
				CountryCode:       "DE",
				Name:              aws.String("Name"),
				Name2:             aws.String("Name2"),
				Name3:             aws.String("Name3"),
				TownArea:          aws.String("TownArea"),
				EmailAddress:      aws.String("foobar@baz.com"),
				PhoneNumber:       aws.String("1234567890"),
				MobilePhoneNumber: aws.String("0987654321"),
				AdditionalAddressLines: []string{
					"AdditionalAddressLine1",
					"AdditionalAddressLine2",
				},
				Reference:     aws.String("Reference"),
				ContactPerson: aws.String("ContactPerson"),
				FaxNumber:     aws.String("0987654321"),
				Gate:          aws.String("Gate"),
				Remarks:       aws.String("Remark"),
				LoadingWindows: ct.LoadingWindows{
					{
						StartsAt: "08:00",
						EndsAt:   "12:00",
					},
				},
			},
			equal: false,
		},
	}
	for idx, test := range tests {
		require.Equal(t, test.equal, test.address1.IsEqual(&test.address2), "Test %d: Expected to be equal: %v and %v", idx, test.address1, test.address2)
	}
}
