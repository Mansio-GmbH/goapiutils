package ct

import (
	"testing"
	"time"

	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/stretchr/testify/require"
)

func TestLoadingWindowApplyTo(t *testing.T) {
	tests := []struct {
		startsAt         string
		endsAt           string
		date             chrono.Date
		expectedStartsAt chrono.Time
		expectedEndsAt   chrono.Time
		expectedError    error
	}{
		{
			startsAt:         "08:00",
			endsAt:           "12:00",
			date:             chrono.NewDate(2024, 1, 1),
			expectedStartsAt: chrono.NewTime(2024, 1, 1, 8, 0, 0, 0, time.UTC),
			expectedEndsAt:   chrono.NewTime(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			startsAt:      "08 00",
			endsAt:        "12:00",
			date:          chrono.NewDate(2024, 1, 1),
			expectedError: ErrorInvalidStartsAtFormat,
		},
		{
			startsAt:      "08:00",
			endsAt:        "12 00",
			date:          chrono.NewDate(2024, 1, 1),
			expectedError: ErrorInvalidEndsAtFormat,
		},
	}

	for _, test := range tests {
		lw := LoadingWindow{
			StartsAt: test.startsAt,
			EndsAt:   test.endsAt,
		}

		loadingWindow, err := lw.ApplyTo(test.date)
		if test.expectedError != nil {
			require.ErrorIs(t, err, test.expectedError)
		} else {
			require.NoError(t, err)
			require.Equal(t, test.expectedStartsAt, loadingWindow.StartsAt)
			require.Equal(t, test.expectedEndsAt, loadingWindow.EndsAt)
		}
	}
}

func TestLoadingWindowIsEqual(t *testing.T) {
	tests := []struct {
		lw1            LoadingWindow
		lw2            LoadingWindow
		expectedResult bool
	}{
		{
			lw1: LoadingWindow{
				StartsAt: "08:00",
				EndsAt:   "12:00",
			},
			lw2: LoadingWindow{
				StartsAt: "08:00",
				EndsAt:   "12:00",
			},
			expectedResult: true,
		},
		{
			lw1: LoadingWindow{
				StartsAt: "08:00",
				EndsAt:   "12:00",
			},

			lw2: LoadingWindow{
				StartsAt: "08:00",
				EndsAt:   "12 00",
			},
			expectedResult: false,
		},
	}

	for _, test := range tests {
		require.Equal(t, test.expectedResult, test.lw1.IsEqual(&test.lw2))
	}
}

func TestLoadingWindowsApplyTo(t *testing.T) {
	tests := []struct {
		lws           LoadingWindows
		date          chrono.Date
		expectedLwsAt LoadingWindowsAtDate
	}{
		{
			lws: LoadingWindows{
				{
					StartsAt: "08:00",
					EndsAt:   "12:00",
				},
				{
					StartsAt: "13:00",
					EndsAt:   "17:00",
				},
			},
			date: chrono.NewDate(2024, 1, 1),
			expectedLwsAt: LoadingWindowsAtDate{
				{
					StartsAt: chrono.NewTime(2024, 1, 1, 8, 0, 0, 0, time.UTC),
					EndsAt:   chrono.NewTime(2024, 1, 1, 12, 0, 0, 0, time.UTC),
				},
				{
					StartsAt: chrono.NewTime(2024, 1, 1, 13, 0, 0, 0, time.UTC),
					EndsAt:   chrono.NewTime(2024, 1, 1, 17, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for _, test := range tests {
		loadingWindows, err := test.lws.ApplyTo(test.date)
		require.NoError(t, err)
		require.Equal(t, len(test.expectedLwsAt), len(loadingWindows))
		for i := range test.expectedLwsAt {
			require.Equal(t, test.expectedLwsAt[i].StartsAt, loadingWindows[i].StartsAt)
			require.Equal(t, test.expectedLwsAt[i].EndsAt, loadingWindows[i].EndsAt)
		}
	}
}

func TestLoadingWindowsIsEqual(t *testing.T) {
	tests := []struct {
		lws1           LoadingWindows
		lws2           LoadingWindows
		expectedResult bool
	}{
		{
			lws1: LoadingWindows{
				{
					StartsAt: "08:00",
					EndsAt:   "12:00",
				},
				{
					StartsAt: "13:00",
					EndsAt:   "17:00",
				},
			},
			lws2: LoadingWindows{
				{
					StartsAt: "08:00",
					EndsAt:   "12:00",
				},
				{
					StartsAt: "13:00",
					EndsAt:   "17:00",
				},
			},
			expectedResult: true,
		},
		{
			lws1: LoadingWindows{
				{
					StartsAt: "08:00",
					EndsAt:   "12:00",
				},
				{
					StartsAt: "13:00",
					EndsAt:   "17:00",
				},
			},
			lws2: LoadingWindows{
				{
					StartsAt: "08:00",
					EndsAt:   "12:00",
				},
				{
					StartsAt: "13:00",
					EndsAt:   "17 00",
				},
			},
			expectedResult: false,
		},
	}

	for _, test := range tests {
		require.Equal(t, test.expectedResult, test.lws1.IsEqual(test.lws2))
	}
}
