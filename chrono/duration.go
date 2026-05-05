package chrono

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Duration int64

type DurationComponents struct {
	Seconds int `json:"seconds" dynamodbav:"seconds"`
	Minutes int `json:"minutes" dynamodbav:"minutes"`
	Hours   int `json:"hours" dynamodbav:"hours"`
	Days    int `json:"days" dynamodbav:"days"`
	Weeks   int `json:"weeks" dynamodbav:"weeks"`
}

func (dc DurationComponents) toDuration() Duration {
	return Second*Duration(dc.Seconds) + Minute*Duration(dc.Minutes) + Hour*Duration(dc.Hours+24*(dc.Days+7*dc.Weeks))
}

func NewDuration(dc DurationComponents) Duration {
	return dc.toDuration()
}

func DurationFrom(d time.Duration) Duration {
	return Duration(d)
}

func (d Duration) String() string {
	return d.ToDurationComponents().String()
}

func (d Duration) StringWithOpts(optFns ...DurationStringerOptFn) string {
	return d.ToDurationComponents().StringWithOpts(optFns...)
}

func (d Duration) Nanoseconds() int64 {
	return d.ToStd().Nanoseconds()
}

func (d Duration) Microseconds() int64 {
	return d.ToStd().Microseconds()
}

func (d Duration) Milliseconds() int64 {
	return d.ToStd().Milliseconds()
}

func (d Duration) Seconds() float64 {
	return d.ToStd().Seconds()
}

func (d Duration) Minutes() float64 {
	return d.ToStd().Minutes()
}

func (d Duration) Hours() float64 {
	return d.ToStd().Hours()
}

func (d Duration) Truncate(m Duration) Duration {
	td := d.ToStd().Truncate(time.Duration(m))
	return Duration(td)
}

func (d Duration) Round(m Duration) Duration {
	rd := d.ToStd().Round(time.Duration(m))
	return Duration(rd)
}

func (d Duration) Abs() Duration {
	ad := d.ToStd().Abs()
	return Duration(ad)
}

func ParseDuration(d string) (Duration, error) {
	val, err := parseDuration(d)
	if err != nil {
		return 0, err
	}
	return Duration(val), nil
}

func (d Duration) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return attributevalue.Marshal(d.ToDurationComponents())
}

func (d *Duration) UnmarshalDynamoDBAttributeValue(v types.AttributeValue) error {
	var (
		dc  DurationComponents
		str string = ""
		err error
	)

	if err = attributevalue.Unmarshal(v, &dc); err == nil {
		*d = dc.toDuration()
		return nil
	}
	if err = attributevalue.Unmarshal(v, &str); err == nil {
		*d, err = parseDuration(str)
		return err
	}
	return errors.New("unable to unmarshal duration")
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ToDurationComponents())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var (
		str string = ""
		err error
	)

	comp := DurationComponents{}
	if err = json.Unmarshal(b, &comp); err == nil {
		*d = comp.toDuration()
		return nil
	}
	if err = json.Unmarshal(b, &str); err == nil {
		*d, err = parseDuration(str)
		return err
	}
	return errors.New("unable to unmarshal duration")
}

func (d Duration) decompose() (weeks, days, hours, minutes, seconds Duration) {
	weeks = d.Truncate(Hour * 24 * 7)
	d = d - weeks
	days = d.Truncate(Hour * 24)
	d = d - days
	hours = d.Truncate(Hour)
	d = d - hours
	minutes = d.Truncate(Minute)
	d = d - minutes
	seconds = d.Truncate(Second)
	return
}

func (d Duration) ToDurationComponents() DurationComponents {
	weeks, days, hours, minutes, seconds := d.decompose()
	return DurationComponents{
		Weeks:   (int)(weeks / (Hour * 24 * 7)),
		Days:    (int)(days / (Hour * 24)),
		Hours:   (int)(hours / Hour),
		Minutes: (int)(minutes / Minute),
		Seconds: (int)(seconds / Second),
	}
}

var (
	iso8601Re = regexp.MustCompile(
		`^P(?:(\d+)W|(?:(\d+)Y)?(?:(\d+)M)?(?:(\d+)D)?(?:T(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?)?)$`,
	)
	compactSegmentRe = regexp.MustCompile(`^(\d+(?:\.\d+)?)([a-zA-Z]+)`)

	unitFactors = map[string]Duration{
		"s": Second, "sec": Second, "secs": Second,
		"second": Second, "seconds": Second,

		"m": Minute, "min": Minute, "mins": Minute,
		"minute": Minute, "minutes": Minute,

		"h": Hour, "hr": Hour, "hrs": Hour,
		"hour": Hour, "hours": Hour,

		"d": Day, "day": Day, "days": Day,

		"w": Week, "wk": Week, "wks": Week,
		"week": Week, "weeks": Week,

		"mo": Day * 30, "month": Day * 30, "months": Day * 30,

		"y": Day * 365, "yr": Day * 365, "yrs": Day * 365,
		"year": Day * 365, "years": Day * 365,
	}
)

func parseISO8601Duration(dstr string) (Duration, error) {
	if dstr == "P" || dstr == "PT" {
		return 0, fmt.Errorf("duration: ISO 8601 must contain at least one component")
	}

	m := iso8601Re.FindStringSubmatch(dstr)
	if len(m) == 0 {
		return 0, errors.New("invalid ISO 8601 duration")
	}

	// Group 1: weeks (mutually exclusive with everything else, per spec).
	if m[1] != "" {
		w, _ := strconv.ParseInt(m[1], 10, 64)
		return Duration(w) * Week, nil
	}

	iso8601Parts := []struct {
		val    string
		factor Duration
	}{
		{m[2], Day * 365},
		{m[3], Day * 30},
		{m[4], Day},
		{m[5], Hour},
		{m[6], Minute},
		{m[7], Second},
	}

	var total Duration
	any := false
	for _, p := range iso8601Parts {
		if p.val == "" {
			continue
		}
		any = true
		n, err := strconv.ParseInt(p.val, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("duration: invalid ISO 8601 number %q: %w", p.val, err)
		}
		total += Duration(n) * p.factor
	}
	if !any {
		return 0, fmt.Errorf("duration: ISO 8601 has no components: %q", dstr)
	}
	return Duration(total), nil
}

func parseCompact(s string) (Duration, error) {
	var total int64
	rest := strings.ReplaceAll(s, " ", "")
	if rest == "" {
		return 0, fmt.Errorf("duration: empty after trimming")
	}

	for len(rest) > 0 {
		m := compactSegmentRe.FindStringSubmatch(rest)
		if m == nil {
			return 0, fmt.Errorf("duration: cannot parse segment near %q in %q", rest, s)
		}

		numStr, unitStr := m[1], strings.ToLower(m[2])
		factor, ok := unitFactors[unitStr]
		if !ok {
			return 0, fmt.Errorf("duration: unknown unit %q", m[2])
		}

		// Allow decimals (e.g. "1.5h") and round to whole seconds.
		f, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return 0, fmt.Errorf("duration: invalid number %q: %w", numStr, err)
		}
		total += int64(f * float64(factor))

		rest = rest[len(m[0]):]
	}

	return Duration(total), nil
}

func parseDuration(dstr string) (Duration, error) {
	dstr = strings.TrimSpace(dstr)
	if isoStr := strings.ToUpper(dstr); isoStr[0] == 'P' {
		return parseISO8601Duration(isoStr)
	}
	return parseCompact(dstr)
}

func (d Duration) IsZero() bool {
	return d == 0
}

func (d Duration) ToStd() time.Duration {
	return time.Duration(d)
}

func (d Duration) ToStdPtr() *time.Duration {
	if d.IsZero() {
		return nil
	}
	std := d.ToStd()
	return &std
}
func (d Duration) Smaller(od Duration) bool {
	return d < od
}

func (d Duration) SmallerOrEqual(od Duration) bool {
	return d <= od
}

func (d Duration) Greater(od Duration) bool {
	return d > od
}

func (d Duration) GreaterOrEqual(od Duration) bool {
	return d >= od
}

type durationComponentLabels struct {
	weeks     string
	days      string
	hours     string
	minutes   string
	seconds   string
	separator string
}

type DurationStringerOptFn func(l *durationComponentLabels)

func WithLabels(weeks, days, hours, minutes, seconds string) DurationStringerOptFn {
	return func(l *durationComponentLabels) {
		l.weeks = weeks
		l.days = days
		l.hours = hours
		l.minutes = minutes
		l.seconds = seconds
	}
}

func WithSeparator(separator string) DurationStringerOptFn {
	return func(l *durationComponentLabels) {
		l.separator = separator
	}
}

func (dc DurationComponents) String() string {
	return dc.StringWithOpts()
}

func (dc DurationComponents) StringWithOpts(optFns ...DurationStringerOptFn) string {
	dcl := durationComponentLabels{
		weeks:     "w",
		days:      "d",
		hours:     "h",
		minutes:   "m",
		seconds:   "s",
		separator: "",
	}
	for _, optFn := range optFns {
		optFn(&dcl)
	}

	components := []string{}
	if dc.Weeks > 0 {
		components = append(components, fmt.Sprintf("%d%s", dc.Weeks, dcl.weeks))
	}
	if dc.Days > 0 {
		components = append(components, fmt.Sprintf("%d%s", dc.Days, dcl.days))
	}
	if dc.Hours > 0 {
		components = append(components, fmt.Sprintf("%d%s", dc.Hours, dcl.hours))
	}
	if dc.Minutes > 0 {
		components = append(components, fmt.Sprintf("%d%s", dc.Minutes, dcl.minutes))
	}
	if dc.Seconds > 0 {
		components = append(components, fmt.Sprintf("%d%s", dc.Seconds, dcl.seconds))
	}
	return strings.Join(components, dcl.separator)
}

func (d Duration) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, "\"%s\"", d.ToDurationComponents().String())
}

func (d *Duration) UnmarshalGQL(v interface{}) error {
	var err error
	switch v := v.(type) {
	case string:
		*d, err = parseDuration(v)
		return err
	case int:
		*d = Duration(v) * Second
		return nil
	case float64:
		*d = Duration(v) * Second
		return nil
	}
	return errors.New("unable to unmarshal duration")
}
