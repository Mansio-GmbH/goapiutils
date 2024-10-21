package ct_test

import (
	"math"
	"testing"

	"github.com/mansio-gmbh/goapiutils/ct"
	"github.com/stretchr/testify/require"
)

func TestTimeDistance_Invalidate(t *testing.T) {
	timeDistance := ct.TimeDistance{
		DurationS: 1000,
		DistanceM: ct.Distance(1000),
	}

	timeDistance.Invalidate()

	require.Equal(t, math.MaxInt, timeDistance.DurationS)
	require.Equal(t, ct.Distance(math.MaxInt64), timeDistance.DistanceM)
	require.Equal(t, ct.ErrInvalidTimeDistance, timeDistance.Problem)
}
