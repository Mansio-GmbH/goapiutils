package chrono

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mansio-gmbh/goapiutils/must"
)

type Time struct {
	val time.Time
}

func (t Time) PtrOrNil() *Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

func (t Time) ToStd() time.Time {
	return t.val
}

func (t Time) ToStdPtr() *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t.val
}

func From(t time.Time) Time {
	return Time{val: t}
}

func FromPtr(t *time.Time) *Time {
	if t == nil {
		return nil
	}
	return &Time{val: *t}
}

func (t Time) After(u Time) bool {
	return t.val.After(u.val)
}

func (t Time) Before(u Time) bool {
	return t.val.Before(u.val)
}

func (t Time) Equal(u Time) bool {
	return t.val.Equal(u.val)
}

func (t Time) AfterDate(u Date) bool {
	return t.val.After(u.val)
}

func (t Time) BeforeDate(u Date) bool {
	return t.val.Before(u.val)
}

func (t Time) EqualDate(u Date) bool {
	return t.val.Equal(u.val)
}

func (t Time) IsZero() bool {
	return t.val.IsZero()
}

func (t Time) Date() Date {
	return Date{
		val: toDate(t.val),
	}
}

func (t Time) Year() int {
	return t.val.Year()
}

func (t Time) Month() time.Month {
	return t.val.Month()
}

func (t Time) Day() int {
	return t.val.Day()
}

func (t Time) Weekday() time.Weekday {
	return t.val.Weekday()
}

func (t Time) Clock() (hour, min, sec int) {
	return t.val.Clock()
}

func (t Time) Hour() int {
	return t.val.Hour()
}

func (t Time) Minute() int {
	return t.val.Minute()
}

func (t Time) Second() int {
	return t.val.Second()
}

func (t Time) Nanosecond() int {
	return t.val.Nanosecond()
}

func (t Time) YearDay() int {
	return t.val.YearDay()
}

func (t Time) Add(d time.Duration) Time {
	return Time{
		val: t.val.Add(d),
	}
}

func (t Time) AddDuration(d Duration) Time {
	return Time{
		val: t.val.Add(d.val),
	}
}

func (t Time) Sub(u Time) Duration {
	return Duration{
		val: t.val.Sub(u.val),
	}
}

func Since(t Time) Duration {
	return Duration{
		val: time.Since(t.val),
	}
}

func Until(t Time) Duration {
	return Duration{
		val: time.Until(t.val),
	}
}

func (t Time) AddDate(years int, months int, days int) Time {
	return Time{
		val: t.val.AddDate(years, months, days),
	}
}

func Now() Time {
	return Time{
		val: time.Now(),
	}
}

func (t Time) UTC() Time {
	return Time{
		val: t.val.UTC(),
	}
}

func (t Time) Local() Time {
	return Time{
		val: t.val.Local(),
	}
}

func (t Time) In(loc *time.Location) Time {
	return Time{
		val: t.val.In(loc),
	}
}

func (t Time) Location() *time.Location {
	return t.val.Location()
}

func (t Time) Zone() (name string, offset int) {
	return t.val.Zone()
}

func (t Time) ZoneBounds() (start, end Time) {
	st, en := t.val.ZoneBounds()
	return Time{
			val: st,
		}, Time{
			val: en,
		}
}

func (t Time) Unix() int64 {
	return t.val.Unix()
}

func (t Time) UnixMilli() int64 {
	return t.val.UnixMilli()
}

func (t Time) UnixMicro() int64 {
	return t.val.UnixMicro()
}

func (t Time) UnixNano() int64 {
	return t.val.UnixNano()
}

func UnixMilli(sec int64) Time {
	return Time{
		val: time.UnixMilli(sec),
	}
}

func UnixMicro(sec int64) Time {
	return Time{
		val: time.UnixMicro(sec),
	}
}

func (t Time) IsDST() bool {
	return t.val.IsDST()
}

func FromUnix(sec int64, nsec int64) Time {
	return Time{
		val: time.Unix(sec, nsec),
	}
}

func NewTime(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return Time{
		val: time.Date(year, month, day, hour, min, sec, nsec, loc),
	}
}

func (t Time) Truncate(d Duration) Time {
	return Time{
		val: t.val.Truncate(d.val),
	}
}

func (t Time) Round(d Duration) Time {
	return Time{
		val: t.val.Round(d.val),
	}
}

type fmtCfg struct {
	layout string
}

func WithLayout(layout string) func(*fmtCfg) {
	return func(cfg *fmtCfg) { cfg.layout = layout }
}

func (t Time) Format(optFns ...func(*fmtCfg)) string {
	cfg := &fmtCfg{layout: time.RFC3339}
	for _, optFn := range optFns {
		optFn(cfg)
	}
	return t.val.Format(cfg.layout)
}

func Parse(str string) (Time, error) {
	val, err := parseTime(str)
	if err != nil {
		return Time{}, err
	}
	return Time{val: val}, nil
}

func ParsePtr(str *string) (Time, error) {
	if str == nil {
		return Time{}, errors.New("input is nil")
	}
	return Parse(*str)
}

func ParseOrNil(str string) *Time {
	t, err := Parse(str)
	if err != nil {
		return nil
	}
	return t.PtrOrNil()
}

func ParsePtrOrNil(str *string) *Time {
	if str != nil {
		return nil
	}
	return ParseOrNil(*str)
}

func MustParse(str string) Time {
	return must.Must(Parse(str))
}

func (t *Time) UnmarshalDynamoDBAttributeValue(v types.AttributeValue) error {
	val := time.Time{}
	if err := attributevalue.Unmarshal(v, &val); err != nil {
		return err
	}
	t.val = val
	return nil
}

func (t Time) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return attributevalue.Marshal(t.val)
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}

func (t *Time) UnmarshalJSON(b []byte) error {
	var (
		str string = ""
		err error
	)

	if err = json.Unmarshal(b, &str); err != nil {
		return err
	}
	if t.val, err = parseTime(str); err != nil {
		return err
	}
	return nil
}

func parseTime(t string) (time.Time, error) {
	if time, err := time.Parse(time.RFC3339, t); err == nil {
		return time, nil
	}
	if time, err := time.Parse("02.01.2006 15:04", t); err == nil {
		return time, nil
	}
	if time, err := time.ParseInLocation("02.01.2006", t, time.Local); err == nil {
		return time, nil
	}
	if time, err := time.ParseInLocation("2006-01-02", t, time.Local); err == nil {
		return time, nil
	}

	return time.Time{}, errors.New("time invalid format")
}
