package timefilter

import (
	"time"

	"github.com/oklog/ulid/v2"
)

func TimeMinMaxSearchValueForULID(t time.Time) string {
	id := ulid.ULID{}
	id.SetTime(ulid.Timestamp(t))
	return id.String()
}
