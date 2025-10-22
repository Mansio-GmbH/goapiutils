package chrono_test

import (
	"encoding/json"
	"testing"

	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/stretchr/testify/require"
)

func TestMarshalTime(t *testing.T) {
	ti := chrono.MustParse("2023-11-29T14:00:00+01:00")
	byt, err := json.Marshal(ti)
	require.NoError(t, err)
	require.Equal(t, `"2023-11-29T14:00:00+01:00"`, string(byt))
}

func TestBeginningOfDay(t *testing.T) {
	tests := []struct {
		input    chrono.Time
		expected chrono.Time
	}{
		{
			input:    chrono.MustParse("2023-11-29T14:00:00+01:00"),
			expected: chrono.MustParse("2023-11-29T00:00:00+01:00"),
		},
		{
			input:    chrono.MustParse("2023-11-29T00:00:00+01:00"),
			expected: chrono.MustParse("2023-11-29T00:00:00+01:00"),
		},
		{
			input:    chrono.MustParse("2023-11-29T23:59:59+01:00"),
			expected: chrono.MustParse("2023-11-29T00:00:00+01:00"),
		},
		{
			input:    chrono.MustParse("2023-11-29T14:00:00Z"),
			expected: chrono.MustParse("2023-11-29T00:00:00Z"),
		},
		{
			input:    chrono.MustParse("2023-11-29T14:12:17-02:00"),
			expected: chrono.MustParse("2023-11-29T00:00:00-02:00"),
		},
	}
	for _, test := range tests {
		result := test.input.BeginningOfDay()
		require.Equal(t, test.expected, result, "Expected "+test.expected.String()+" but got "+result.String()+" for input "+test.input.String())
	}
}

func TestEndOfDay(t *testing.T) {
	tests := []struct {
		input    chrono.Time
		expected chrono.Time
	}{
		{
			input:    chrono.MustParse("2023-11-29T14:00:00+01:00"),
			expected: chrono.MustParse("2023-11-29T23:59:59+01:00"),
		},
		{
			input:    chrono.MustParse("2023-11-29T00:00:00+01:00"),
			expected: chrono.MustParse("2023-11-29T23:59:59+01:00"),
		},
		{
			input:    chrono.MustParse("2023-11-29T23:59:59+01:00"),
			expected: chrono.MustParse("2023-11-29T23:59:59+01:00"),
		},
		{
			input:    chrono.MustParse("2023-11-29T14:00:00Z"),
			expected: chrono.MustParse("2023-11-29T23:59:59Z"),
		},
	}
	for _, test := range tests {
		result := test.input.EndOfDay()
		require.Equal(t, test.expected, result, "Expected "+test.expected.String()+" but got "+result.String()+" for input "+test.input.String())
	}
}

func TestBeginningOfMonth(t *testing.T) {
	tests := []struct {
		input    chrono.Time
		expected chrono.Time
	}{
		{
			input:    chrono.MustParse("2023-11-29T14:00:00+01:00"),
			expected: chrono.MustParse("2023-11-01T00:00:00+01:00"),
		},
		{
			input:    chrono.MustParse("2023-11-01T14:00:00+01:00"),
			expected: chrono.MustParse("2023-11-01T00:00:00+01:00"),
		},
	}
	for _, test := range tests {
		result := test.input.BeginningOfMonth()
		require.Equal(t, test.expected, result, "Expected "+test.expected.String()+" but got "+result.String()+" for input "+test.input.String())
	}
}

func TestEndOfMonth(t *testing.T) {
	tests := []struct {
		input    chrono.Time
		expected chrono.Time
	}{
		{
			input:    chrono.MustParse("2023-11-29T14:00:00+01:00"),
			expected: chrono.MustParse("2023-11-30T23:59:59+01:00"),
		},
		{
			input:    chrono.MustParse("2023-11-01T14:00:00+01:00"),
			expected: chrono.MustParse("2023-11-30T23:59:59+01:00"),
		},
		{
			input:    chrono.MustParse("2023-05-14T14:00:00+01:00"),
			expected: chrono.MustParse("2023-05-31T23:59:59+01:00"),
		},
		{
			input:    chrono.MustParse("2023-02-14T14:00:00+01:00"),
			expected: chrono.MustParse("2023-02-28T23:59:59+01:00"),
		},
		{
			input:    chrono.MustParse("2024-02-14T14:00:00+01:00"),
			expected: chrono.MustParse("2024-02-29T23:59:59+01:00"),
		},
	}
	for _, test := range tests {
		result := test.input.EndOfMonth()
		require.Equal(t, test.expected, result, "Expected "+test.expected.String()+" but got "+result.String()+" for input "+test.input.String())
	}
}

func TestBeginningOfYear(t *testing.T) {
	tests := []struct {
		input    chrono.Time
		expected chrono.Time
	}{
		{
			input:    chrono.MustParse("2023-11-29T14:00:00+01:00"),
			expected: chrono.MustParse("2023-01-01T00:00:00+01:00"),
		},
		{
			input:    chrono.MustParse("2023-01-01T14:00:00+01:00"),
			expected: chrono.MustParse("2023-01-01T00:00:00+01:00"),
		},
	}
	for _, test := range tests {
		result := test.input.BeginningOfYear()
		require.Equal(t, test.expected, result, "Expected "+test.expected.String()+" but got "+result.String()+" for input "+test.input.String())
	}
}

func TestEndOfYear(t *testing.T) {
	tests := []struct {
		input    chrono.Time
		expected chrono.Time
	}{
		{
			input:    chrono.MustParse("2023-11-29T14:00:00+01:00"),
			expected: chrono.MustParse("2023-12-31T23:59:59+01:00"),
		},
		{
			input:    chrono.MustParse("2023-01-01T14:00:00+01:00"),
			expected: chrono.MustParse("2023-12-31T23:59:59+01:00"),
		},
	}
	for _, test := range tests {
		result := test.input.EndOfYear()
		require.Equal(t, test.expected, result, "Expected "+test.expected.String()+" but got "+result.String()+" for input "+test.input.String())
	}
}

func TestBeginningOfWeek(t *testing.T) {
	tests := []struct {
		input    chrono.Time
		expected chrono.Time
	}{
		{
			input:    chrono.MustParse("2023-11-29T14:00:00+01:00"),
			expected: chrono.MustParse("2023-11-27T00:00:00+01:00"),
		},
		{
			input:    chrono.MustParse("2023-11-27T14:00:00+01:00"),
			expected: chrono.MustParse("2023-11-27T00:00:00+01:00"),
		},
	}
	for _, test := range tests {
		result := test.input.BeginningOfWeek()
		require.Equal(t, test.expected, result, "Expected "+test.expected.String()+" but got "+result.String()+" for input "+test.input.String())
	}
}

func TestEndOfWeek(t *testing.T) {
	tests := []struct {
		input    chrono.Time
		expected chrono.Time
	}{
		{
			input:    chrono.MustParse("2023-11-29T14:00:00+01:00"),
			expected: chrono.MustParse("2023-12-03T23:59:59+01:00"),
		},
		{
			input:    chrono.MustParse("2023-11-27T14:00:00+01:00"),
			expected: chrono.MustParse("2023-12-03T23:59:59+01:00"),
		},
	}
	for _, test := range tests {
		result := test.input.EndOfWeek()
		require.Equal(t, test.expected, result, "Expected "+test.expected.String()+" but got "+result.String()+" for input "+test.input.String())
	}
}

func TestTimeIsZero(t *testing.T) {
	parsedTime, err := chrono.Parse("0001-01-01T00:00:00Z")
	require.NoError(t, err)
	require.True(t, parsedTime.IsZero())
	require.True(t, chrono.Time{}.IsZero())
}

func TestTimeFuture(t *testing.T) {
	futureTime := chrono.Now().AddDate(0, 0, 1)
	require.True(t, futureTime.Future())
	require.False(t, futureTime.Past())
}

func TestTimePast(t *testing.T) {
	pastTime := chrono.Now().AddDate(0, 0, -1)
	require.True(t, pastTime.Past())
	require.False(t, pastTime.Future())
}
