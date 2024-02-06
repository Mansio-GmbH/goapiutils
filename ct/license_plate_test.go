package ct

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLicensePlate_MarshalJSON(t *testing.T) {
	lp := LicensePlate{
		licensePlateSegments: []string{"K", "AM", "123"},
	}
	val, err := lp.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, "K-AM-123", string(val))
}

func TestLicensePlate_UnmarshalJSON(t *testing.T) {
	lp := &LicensePlate{}
	err := lp.UnmarshalJSON([]byte(`"K-AM 123"`))
	require.NoError(t, err)
	require.Equal(t, []string{"K", "AM", "123"}, lp.licensePlateSegments)
}

func TestLicensePlate_MarshalDynamoDBAttributeValue(t *testing.T) {
	lp := LicensePlate{
		licensePlateSegments: []string{"K", "AM", "123"},
	}
	lpav, err := lp.MarshalDynamoDBAttributeValue()
	require.NoError(t, err)

	var unmarshaledLP LicensePlate
	err = unmarshaledLP.UnmarshalDynamoDBAttributeValue(lpav)
	require.NoError(t, err)
	require.Equal(t, lp, unmarshaledLP)
}
