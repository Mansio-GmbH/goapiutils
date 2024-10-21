package ct

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
)

var ErrInvalidTimeDistance = errors.New("invalid time distance")

type TimeDistance struct {
	DurationS int
	DistanceM Distance
	Problem   error
}

// Invalidate sets the TimeDistance to invalid and sets the duration and distance to the maximum possible value.
// The maximum values will prevent tools to produce wrong results.
func (t *TimeDistance) Invalidate() {
	t.Problem = ErrInvalidTimeDistance
	t.DurationS = math.MaxInt
	t.DistanceM = Distance(math.MaxInt64)
}

func (t *TimeDistance) MarshalJSON() ([]byte, error) {
	start := `{"duration_s":` + strconv.Itoa(t.DurationS) + `,"distance_m":` + strconv.FormatInt(int64(t.DistanceM), 10)
	if t.Problem != nil {
		return []byte(start + `",problem":"` + t.Problem.Error() + `"}`), nil
	}
	return []byte(start + `}`), nil
}

func (t *TimeDistance) UnmarshalJSON(b []byte) error {
	tmp := struct {
		DurationS int      `json:"duration_s"`
		DistanceM Distance `json:"distance_m"`
		Problem   string   `json:"problem,omitempty"`
	}{}

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
