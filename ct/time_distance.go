package ct

import (
	"errors"
	"math"
)

var ErrInvalidTimeDistance = errors.New("invalid time distance")

type TimeDistance struct {
	DurationS int      `json:"duration_s"`
	DistanceM Distance `json:"distance_m"`
	Problem   error    `json:"problem,omitempty"`
}

// Invalidate sets the TimeDistance to invalid and sets the duration and distance to the maximum possible value.
// The maximum values will prevent tools to produce wrong results.
func (t *TimeDistance) Invalidate() {
	t.Problem = ErrInvalidTimeDistance
	t.DurationS = math.MaxInt
	t.DistanceM = Distance(math.MaxInt64)
}
