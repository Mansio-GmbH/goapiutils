package multikey_test

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mansio-gmbh/goapiutils/multikey"
	"github.com/mansio-gmbh/goapiutils/must"
	"github.com/mansio-gmbh/goapiutils/network"
	"github.com/mansio-gmbh/goapiutils/tenant"
	"github.com/stretchr/testify/require"
)

type testType struct{}

func (testType) ToString() string {
	return "BLUBB"
}

func TestKey(t *testing.T) {
	testcases := []struct {
		parts    []any
		expected string
	}{
		{
			parts: []any{
				multikey.TENANT,
				1234,
				must.Must(time.Parse(time.RFC3339, "2023-10-26T10:00:00+02:00")),
				"foo",
			},
			expected: "TENANT#1234#2023-10-26T08:00:00Z#foo",
		},
		{
			parts: []any{
				testType{},
			},
			expected: "BLUBB",
		},
		{
			parts: []any{
				tenant.Tenant{ID: "test ~tenant"},
				network.Network{ID: "test-network"},
			},
			expected: "TENANT#TESTTENANT#NETWORK#TESTNETWORK",
		},
		{
			parts: []any{
				&tenant.Tenant{ID: "test ~tenant"},
				&network.Network{ID: "test-network"},
			},
			expected: "TENANT#TESTTENANT#NETWORK#TESTNETWORK",
		},
		{
			parts: []any{
				aws.String("foo"),
			},
			expected: "foo",
		},
	}

	for _, test := range testcases {
		key := *multikey.Key(test.parts[0], test.parts[1:]...)
		require.Equal(t, test.expected, key)
	}
}

func TestKeyAW(t *testing.T) {
	key := multikey.KeyAV(tenant.Tenant{ID: "foo"}, multikey.ADDRESS)
	require.Equal(t, types.AttributeValueMemberS{Value: "TENANT#FOO#ADDRESS"}, *key.(*types.AttributeValueMemberS))
}
