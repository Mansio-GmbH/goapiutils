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
