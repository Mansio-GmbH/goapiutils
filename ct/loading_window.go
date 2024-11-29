package ct

import (
	"errors"
	"fmt"

	"github.com/mansio-gmbh/goapiutils/chrono"
)

type LoadingWindow struct {
	StartsAt string `json:"startsAt" dynamodbav:"startsAt"`
	EndsAt   string `json:"endsAt" dynamodbav:"endsAt"`
}

type LoadingWindows []LoadingWindow

type LoadingWindowAtDate struct {
	StartsAt chrono.Time `json:"startsAt" dynamodbav:"startsAt"`
	EndsAt   chrono.Time `json:"endsAt" dynamodbav:"endsAt"`
}

type LoadingWindowsAtDate []LoadingWindowAtDate

var (
	ErrorInvalidStartsAtFormat = errors.New("invalid startsAt format")
	ErrorInvalidEndsAtFormat   = errors.New("invalid endsAt format")
)

func (lw LoadingWindows) IsEqual(other LoadingWindows) bool {
	if len(lw) != len(other) {
		return false
	}
	for i, w := range lw {
		if !w.IsEqual(other[i]) {
			return false
		}
	}
	return true
}

func (lw LoadingWindows) ApplyTo(date chrono.Date) (loadingWindows LoadingWindowsAtDate, err error) {
	loadingWindows = make(LoadingWindowsAtDate, len(lw))
	for i, w := range lw {
		loadingWindow, err := w.ApplyTo(date)
		if err != nil {
			return nil, err
		}
		loadingWindows[i] = loadingWindow
	}
	return
}

func (lw LoadingWindow) IsEqual(other LoadingWindow) bool {
	return lw.StartsAt == other.StartsAt && lw.EndsAt == other.EndsAt
}

func (lw LoadingWindow) ApplyTo(date chrono.Date) (loadingWindow LoadingWindowAtDate, err error) {
	startsAtHour, startsAtMinute, endsAtHour, endsAtMinute, err := lw.parse()
	if err != nil {
		return
	}
	loadingWindow.StartsAt = chrono.NewTime(date.Year(), date.Month(), date.Day(), startsAtHour, startsAtMinute, 0, 0, date.Location())
	loadingWindow.EndsAt = chrono.NewTime(date.Year(), date.Month(), date.Day(), endsAtHour, endsAtMinute, 0, 0, date.Location())
	return
}

func (lw LoadingWindow) parse() (startsAtHour, startsAtMinute, endsAtHour, endsAtMinute int, err error) {
	// The format is "HH:MM"
	_, err = fmt.Sscanf(lw.StartsAt, "%d:%d", &startsAtHour, &startsAtMinute)
	if err != nil {
		err = ErrorInvalidStartsAtFormat
		return
	}
	if _, err = fmt.Sscanf(lw.EndsAt, "%d:%d", &endsAtHour, &endsAtMinute); err != nil {
		err = ErrorInvalidEndsAtFormat
		return
	}
	return
}
