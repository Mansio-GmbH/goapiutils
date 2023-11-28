package chrono

import "time"

type Month = time.Month

const (
	January   = time.January
	February  = time.February
	March     = time.March
	April     = time.April
	May       = time.May
	June      = time.June
	July      = time.July
	August    = time.August
	September = time.September
	October   = time.October
	November  = time.November
	December  = time.December
)

type Weekday = time.Weekday

const (
	Sunday    = time.Sunday
	Monday    = time.Monday
	Tuesday   = time.Tuesday
	Wednesday = time.Wednesday
	Thursday  = time.Thursday
	Friday    = time.Friday
	Saturday  = time.Saturday
)

const (
	Nanosecond  = time.Nanosecond
	Microsecond = time.Microsecond
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	Hour        = time.Hour
	Day         = Hour * 24
	Week        = 7 * Day
)

type Location = time.Location
