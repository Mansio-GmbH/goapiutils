package chrono

import (
	"encoding/json"
	"errors"
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
	return time.Duration(d).String()
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
	return attributevalue.Marshal(d.toDurationComponents())
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
	return json.Marshal(d.toDurationComponents())
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

func (d Duration) Decompose() (weeks, days, hours, minutes, seconds Duration) {
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

func (d Duration) toDurationComponents() DurationComponents {
	weeks, days, hours, minutes, seconds := d.Decompose()
	return DurationComponents{
		Weeks:   (int)(weeks / (Hour * 24 * 7)),
		Days:    (int)(days / (Hour * 24)),
		Hours:   (int)(hours / Hour),
		Minutes: (int)(minutes / Minute),
		Seconds: (int)(seconds / Second),
	}
}

func parseDuration(dstr string) (Duration, error) {
	d, err := time.ParseDuration(dstr)
	return Duration(d), err
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
