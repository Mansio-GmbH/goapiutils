package chrono_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/stretchr/testify/require"
)

func TestNewDate(t *testing.T) {
	ti := chrono.NewDate(2023, 11, 29)
	require.Equal(t, "2023-11-29", ti.Format(chrono.WithLayout("2006-01-02")))
}

func TestMarshalDate(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	byt, err := json.Marshal(ti)
	require.NoError(t, err)
	require.Equal(t, `"2023-11-29"`, string(byt))
}

func TestAddDate(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	newDate := ti.AddDate(1, 0, 0)
	require.Equal(t, "2024-11-29", newDate.Format(chrono.WithLayout("2006-01-02")))

	newDate = ti.AddDate(0, 1, 0)
	require.Equal(t, "2023-12-29", newDate.Format(chrono.WithLayout("2006-01-02")))

	newDate = ti.AddDate(0, 0, 1)
	require.Equal(t, "2023-11-30", newDate.Format(chrono.WithLayout("2006-01-02")))
}

func TestSubDate(t *testing.T) {
	ti1 := chrono.MustParseDate("2023-11-29")
	ti2 := chrono.MustParseDate("2023-11-28")
	days := ti1.SubDate(ti2)
	require.Equal(t, 1, days)

	ti3 := chrono.MustParseDate("2023-11-30")
	days = ti1.SubDate(ti3)
	require.Equal(t, -1, days)

	ti4 := chrono.MustParseDate("2023-11-29")
	days = ti1.SubDate(ti4)
	require.Equal(t, 0, days)

	ti5 := chrono.MustParseDate("2024-11-29")
	days = ti5.SubDate(ti1)
	require.Equal(t, 366, days)
}

func TestDateBeginningOfMonth(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	beginning := ti.BeginningOfMonth()
	require.Equal(t, "2023-11-01", beginning.Format(chrono.WithLayout("2006-01-02")))

	ti = chrono.MustParseDate("2024-02-29")
	beginning = ti.BeginningOfMonth()
	require.Equal(t, "2024-02-01", beginning.Format(chrono.WithLayout("2006-01-02")))
}

func TestDateEndOfMonth(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	end := ti.EndOfMonth()
	require.Equal(t, "2023-11-30", end.Format(chrono.WithLayout("2006-01-02")))

	ti = chrono.MustParseDate("2024-02-12")
	end = ti.EndOfMonth()
	require.Equal(t, "2024-02-29", end.Format(chrono.WithLayout("2006-01-02")))
}

func TestDateBeginningOfYear(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	beginning := ti.BeginningOfYear()
	require.Equal(t, "2023-01-01", beginning.Format(chrono.WithLayout("2006-01-02")))

	ti = chrono.MustParseDate("2024-02-29")
	beginning = ti.BeginningOfYear()
	require.Equal(t, "2024-01-01", beginning.Format(chrono.WithLayout("2006-01-02")))
}

func TestDateEndOfYear(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	end := ti.EndOfYear()
	require.Equal(t, "2023-12-31", end.Format(chrono.WithLayout("2006-01-02")))

	ti = chrono.MustParseDate("2024-02-29")
	end = ti.EndOfYear()
	require.Equal(t, "2024-12-31", end.Format(chrono.WithLayout("2006-01-02")))
}

func TestEqual(t *testing.T) {
	ti1 := chrono.MustParseDate("2023-11-29")
	ti2 := chrono.MustParseDate("2023-11-29")
	ti3 := chrono.MustParseDate("2023-11-30")

	require.True(t, ti1.EqualDate(ti2))
	require.True(t, ti1.Equal(chrono.MustParse("2023-11-29T12:12:34Z")))
	require.False(t, ti1.EqualDate(ti3))
	require.False(t, ti2.EqualDate(ti3))
	require.True(t, ti1.BeforeDate(ti3))
	require.False(t, ti1.BeforeDate(ti2))
	require.False(t, ti2.BeforeDate(ti1))
}

func TestAfter(t *testing.T) {
	ti1 := chrono.MustParseDate("2023-11-29")
	ti2 := chrono.MustParseDate("2023-11-30")
	ti3 := chrono.MustParseDate("2023-11-28")

	require.True(t, ti2.AfterDate(ti1))
	require.False(t, ti1.AfterDate(ti2))
	require.True(t, ti1.AfterDate(ti3))
	require.False(t, ti3.AfterDate(ti1))
	require.True(t, ti1.After(ti3.Time()))
}

func TestBefore(t *testing.T) {
	ti1 := chrono.MustParseDate("2023-11-29")
	ti2 := chrono.MustParseDate("2023-11-30")
	ti3 := chrono.MustParseDate("2023-11-28")

	require.True(t, ti1.BeforeDate(ti2))
	require.False(t, ti2.BeforeDate(ti1))
	require.True(t, ti3.BeforeDate(ti1))
	require.False(t, ti1.BeforeDate(ti3))
	require.True(t, ti3.Before(ti1.Time()))
}

func TestIsZero(t *testing.T) {
	ti := chrono.Date{}
	require.True(t, ti.IsZero())

	ti = chrono.MustParseDate("2023-11-29")
	require.False(t, ti.IsZero())
}

func TestYear(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	require.Equal(t, 2023, ti.Year())

	ti = chrono.MustParseDate("2024-02-29")
	require.Equal(t, 2024, ti.Year())
}

func TestMonth(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	require.Equal(t, time.November, ti.Month())

	ti = chrono.MustParseDate("2024-02-29")
	require.Equal(t, time.February, ti.Month())
}

func TestDay(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	require.Equal(t, 29, ti.Day())

	ti = chrono.MustParseDate("2024-02-29")
	require.Equal(t, 29, ti.Day())
}

func TestWeekDay(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	require.Equal(t, time.Wednesday, ti.Weekday())

	ti = chrono.MustParseDate("2024-02-29")
	require.Equal(t, time.Thursday, ti.Weekday())
}

func TestYearDay(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	require.Equal(t, 333, ti.YearDay())

	ti = chrono.MustParseDate("2024-02-29")
	require.Equal(t, 60, ti.YearDay())
}

func TestDateFromTimePtr(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	timePtr := ti.Time().PtrOrNil()
	require.NotNil(t, timePtr)

	date := chrono.DateFromTimePtr(timePtr)
	require.Equal(t, ti, *date)

	var nilTime *chrono.Time
	date = chrono.DateFromTimePtr(nilTime)
	require.Nil(t, date)
}

func TestPast(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	require.True(t, ti.Past())
}

func TestDayBefore(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	dayBefore := ti.DayBefore()
	require.Equal(t, "2023-11-28", dayBefore.Format(chrono.WithLayout("2006-01-02")))

	ti = chrono.MustParseDate("2024-03-01")
	dayBefore = ti.DayBefore()
	require.Equal(t, "2024-02-29", dayBefore.Format(chrono.WithLayout("2006-01-02")))
}

func TestDayAfter(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	dayAfter := ti.DayAfter()
	require.Equal(t, "2023-11-30", dayAfter.Format(chrono.WithLayout("2006-01-02")))

	ti = chrono.MustParseDate("2024-02-28")
	dayAfter = ti.DayAfter()
	require.Equal(t, "2024-02-29", dayAfter.Format(chrono.WithLayout("2006-01-02")))
}

func TestYesterday(t *testing.T) {
	require.Equal(t, chrono.Today().DayBefore(), chrono.Yesterday())
}
