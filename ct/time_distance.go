package ct

import (
	"encoding/json"
	"errors"
	"math"
)

var ErrInvalidTimeDistance = errors.New("invalid time distance")

type TimeDistance struct {
	DurationS int
	DistanceM Distance
	Problem   error
}

type timeDistanceJson struct {
	DurationS int      `json:"durationS"`
	DistanceM Distance `json:"distanceM"`
	Problem   string   `json:"problem,omitempty"`
}

// Invalidate sets the TimeDistance to invalid and sets the duration and distance to the maximum possible value.
// The maximum values will prevent tools to produce wrong results.
func (t *TimeDistance) Invalidate() {
	t.Problem = ErrInvalidTimeDistance
	t.DurationS = math.MaxInt
	t.DistanceM = Distance(math.MaxInt64)
}

func (t *TimeDistance) MarshalJSON() ([]byte, error) {
	tmp := timeDistanceJson{
		DurationS: t.DurationS,
		DistanceM: t.DistanceM,
	}
	if t.Problem != nil {
		tmp.Problem = t.Problem.Error()
	}

	return json.Marshal(tmp)
}

func (t *TimeDistance) UnmarshalJSON(b []byte) error {
	var tmp timeDistanceJson
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return errors.New("unable to unmarshal time distance")
	}

	t.DurationS = tmp.DurationS
	t.DistanceM = tmp.DistanceM
	if tmp.Problem != "" {
		if tmp.Problem == ErrInvalidTimeDistance.Error() {
			t.Problem = ErrInvalidTimeDistance
		} else {
			t.Problem = errors.New(tmp.Problem)
		}
	}
	return nil
}
