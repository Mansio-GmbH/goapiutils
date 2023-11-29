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

func NewDate(year int, month time.Month, day int, loc *time.Location) Date {
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

func (t *Date) Time() Time {
	return Time{val: t.val}
}

func (d Date) After(u Time) bool {
	return d.val.After(u.val)
}

func (d Date) Before(u Time) bool {
	return d.val.Before(u.val)
}

func (d Date) Equal(u Time) bool {
	return d.val.Equal(u.val)
}

func (d Date) AfterDate(u Date) bool {
	return d.val.After(u.val)
}

func (d Date) BeforeDate(u Date) bool {
	return d.val.Before(u.val)
}

func (d Date) EqualDate(u Date) bool {
	return d.val.Equal(u.val)
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

func Today() Date {
	return Date{
		val: toDate(time.Now()),
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

func (d Date) In(loc *time.Location) Date {
	return Date{
		val: d.val.In(loc),
	}
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
	if str != nil {
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

func (t *Date) UnmarshalDynamoDBAttributeValue(v types.AttributeValue) error {
	val := time.Time{}
	if err := attributevalue.Unmarshal(v, &val); err != nil {
		return err
	}
	t.val = val
	return nil
}

func (t Date) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return attributevalue.Marshal(t.val)
}

func (t Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}

func (t *Date) UnmarshalJSON(b []byte) error {
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
	t.val = toDate(val)
	return nil
}

func toDate(ti time.Time) time.Time {
	year, month, day := ti.Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, ti.Location())
	return date
}
