package chrono_test

import (
	"encoding/json"
	"testing"

	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/stretchr/testify/require"
)

func TestMarshalDate(t *testing.T) {
	ti := chrono.MustParseDate("2023-11-29")
	byt, err := json.Marshal(ti)
	require.NoError(t, err)
	require.Equal(t, `"2023-11-29"`, string(byt))
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
