package chrono

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mansio-gmbh/goapiutils/must"
)

type Date struct {
	val time.Time
}

func NewDate(year int, month time.Month, day int) Date {
	return NewDateWithLocation(year, month, day, nil)
}

func NewDateWithLocation(year int, month time.Month, day int, loc *time.Location) Date {
	if loc == nil {
		loc = time.UTC
	}
	return Date{
		val: time.Date(year, month, day, 0, 0, 0, 0, loc),
	}
}

func DateFromTime(t Time) Date {
	return t.Date()
}

func DateFromTimePtr(t *Time) *Date {
	if t == nil {
		return nil
	}
	d := t.Date()
	return &d
}

func (d Date) PtrOrNil() *Date {
	if d.IsZero() {
		return nil
	}
	return &d
}

func (d *Date) Time() Time {
	return Time{val: d.val}
}

func (d Date) After(u Time) bool {
	return d.val.After(u.val)
}

func (d Date) Before(u Time) bool {
	return d.BeforeDate(u.Date())
}

func (d Date) Equal(u Time) bool {
	return d.EqualDate(u.Date())
}

func (d Date) AfterDate(u Date) bool {
	dDay := d.val.Day()
	dMonth := d.val.Month()
	dYear := d.val.Year()
	uDay := u.val.Day()
	uMonth := u.val.Month()
	uYear := u.val.Year()
	return dYear > uYear || (dYear == uYear && (dMonth > uMonth || (dMonth == uMonth && dDay > uDay)))
}

func (d Date) BeforeDate(u Date) bool {
	dDay := d.val.Day()
	dMonth := d.val.Month()
	dYear := d.val.Year()
	uDay := u.val.Day()
	uMonth := u.val.Month()
	uYear := u.val.Year()
	return dYear < uYear || (dYear == uYear && (dMonth < uMonth || (dMonth == uMonth && dDay < uDay)))
}

func (d Date) EqualDate(u Date) bool {
	return d.Day() == u.Day() &&
		d.Month() == u.Month() &&
		d.Year() == u.Year()
}

func (d Date) Past() bool {
	return d.BeforeDate(Today())
}

func (d Date) Future() bool {
	return d.AfterDate(Today())
}

func (d Date) IsZero() bool {
	return d.val.IsZero()
}

func (d Date) Year() int {
	return d.val.Year()
}

func (d Date) Month() time.Month {
	return d.val.Month()
}

func (d Date) Day() int {
	return d.val.Day()
}

func (d Date) Weekday() time.Weekday {
	return d.val.Weekday()
}

func (d Date) YearDay() int {
	return d.val.YearDay()
}

func (d Date) AddDate(years int, months int, days int) Date {
	return Date{
		val: d.val.AddDate(years, months, days),
	}
}

func (d Date) SubDate(u Date) (days int) {
	return int(d.val.Sub(u.val).Hours() / 24)
}

func (d Date) DayBefore() Date {
	return Date{
		val: d.val.AddDate(0, 0, -1),
	}
}

func (d Date) DayAfter() Date {
	return Date{
		val: d.val.AddDate(0, 0, 1),
	}
}

func (d Date) BeginningOfMonth() Date {
	return NewDateWithLocation(d.Year(), d.Month(), 1, d.val.Location())
}

func (d Date) EndOfMonth() Date {
	return NewDateWithLocation(d.Year(), d.Month()+1, 0, d.val.Location())
}

func (d Date) BeginningOfYear() Date {
	return NewDateWithLocation(d.Year(), time.January, 1, d.val.Location())
}

func (d Date) EndOfYear() Date {
	return NewDateWithLocation(d.Year(), time.December, 31, d.val.Location())
}

func Today() Date {
	return Date{
		val: toDate(time.Now()),
	}
}

func Yesterday() Date {
	return Date{
		val: toDate(time.Now().AddDate(0, 0, -1)),
	}
}

func DateFrom(time time.Time) Date {
	return Date{
		val: toDate(time),
	}
}

func DateFromPtr(t *time.Time) *Date {
	if t == nil {
		return nil
	}
	return &Date{
		val: toDate(*t),
	}
}

func (d Date) UTC() Date {
	return Date{
		val: d.val.UTC(),
	}
}

func (d Date) Local() Date {
	return Date{
		val: d.val.Local(),
	}
}

// In returns a new Date in the provided location with same year, month, day values.
func (d Date) In(loc *time.Location) Date {
	return NewDateWithLocation(d.Year(), d.Month(), d.Day(), loc)
}

func (d Date) Location() *time.Location {
	return d.val.Location()
}

func (d Date) Unix() int64 {
	return d.val.Unix()
}

func (d Date) UnixMilli() int64 {
	return d.val.UnixMilli()
}

func (d Date) UnixMicro() int64 {
	return d.val.UnixMicro()
}

func (d Date) UnixNano() int64 {
	return d.val.UnixNano()
}

func (d Date) ToStd() time.Time {
	return d.val
}

func (d Date) ToStdPtr() *time.Time {
	if d.IsZero() {
		return nil
	}
	t := d.val
	return &t
}

func ParseDate(str string) (Date, error) {
	t, err := parseTime(str)
	if err != nil {
		return Date{}, err
	}
	return Time{val: t}.Date(), nil
}

func ParseDatePtr(str *string) (Date, error) {
	if str == nil {
		return Date{}, errors.New("input is nil")
	}
	return ParseDate(*str)
}

func ParseDateOrNil(str string) *Date {
	t, err := parseTime(str)
	if err != nil {
		return nil
	}
	return Time{val: t}.Date().PtrOrNil()
}

func ParseDatePtrOrNil(str *string) *Date {
	if str == nil {
		return nil
	}
	return ParseDateOrNil(*str)
}

func MustParseDate(str string) Date {
	return must.Must(ParseDate(str))
}

func (d Date) Format(optFns ...func(*fmtCfg)) string {
	cfg := &fmtCfg{layout: "2006-01-02"}
	for _, optFn := range optFns {
		optFn(cfg)
	}
	return d.val.Format(cfg.layout)
}

func (d *Date) UnmarshalDynamoDBAttributeValue(v types.AttributeValue) error {
	val := time.Time{}
	if err := attributevalue.Unmarshal(v, &val); err != nil {
		return err
	}
	d.val = val
	return nil
}

func (d Date) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return attributevalue.Marshal(d.val)
}

func (d Date) WithTime(hours, minutes int) Time {
	return Time{
		val: time.Date(d.Year(), d.Month(), d.Day(), hours, minutes, 0, 0, d.Location()),
	}
}

func (d Date) MarshalJSON() ([]byte, error) {
	dateStr := d.val.Format("2006-01-02")
	return json.Marshal(dateStr)
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var (
		str string = ""
		err error
	)

	if err = json.Unmarshal(b, &str); err != nil {
		return err
	}
	val, err := parseTime(str)
	if err != nil {
		return err
	}
	d.val = toDate(val)
	return nil
}

func toDate(ti time.Time) time.Time {
	year, month, day := ti.Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, ti.Location())
	return date
}
