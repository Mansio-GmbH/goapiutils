package timefilter

import (
	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/oklog/ulid/v2"
)

func TimeMinMaxSearchValueForULID(t chrono.Time) string {
	id := ulid.ULID{}
	id.SetTime(ulid.Timestamp(t.ToStd()))
	return id.String()
}
