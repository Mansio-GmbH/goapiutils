package chrono

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Duration struct {
	val time.Duration
}

func (d Duration) String() string {
	return d.val.String()
}

func (d Duration) Nanoseconds() int64 {
	return d.val.Nanoseconds()
}

func (d Duration) Microseconds() int64 {
	return d.val.Microseconds()
}

func (d Duration) Milliseconds() int64 {
	return d.val.Milliseconds()
}

func (d Duration) Seconds() float64 {
	return d.val.Seconds()
}

func (d Duration) Minutes() float64 {
	return d.val.Minutes()
}

func (d Duration) Hours() float64 {
	return d.val.Hours()
}

func (d Duration) Truncate(m Duration) Duration {
	return Duration{val: d.val.Truncate(m.val)}
}

func (d Duration) Round(m Duration) Duration {
	return Duration{val: d.val.Round(m.val)}
}

func (d Duration) Abs() Duration {
	return Duration{val: d.val.Abs()}
}

func (d Duration) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	str := d.val.String()
	return attributevalue.Marshal(str)
}

func (d *Duration) UnmarshalDynamoDBAttributeValue(v types.AttributeValue) error {
	var (
		str string = ""
		err error
	)

	if err := attributevalue.Unmarshal(v, &str); err != nil {
		return err
	}
	if d.val, err = parseDuration(str); err != nil {
		return err
	}
	return nil
}

func (d *Duration) MarshalJSON() ([]byte, error) {
	str := d.val.String()
	return json.Marshal(str)
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var (
		str string = ""
		err error
	)

	if err = json.Unmarshal(b, &str); err != nil {
		return err
	}
	if d.val, err = parseDuration(str); err != nil {
		return err
	}
	return nil
}

func parseDuration(d string) (time.Duration, error) {
	return time.ParseDuration(d)
}
