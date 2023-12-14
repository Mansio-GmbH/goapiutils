package chrono_test

import (
	"testing"
	"time"

	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/stretchr/testify/assert"
)

func TestDurationComparison(t *testing.T) {

	d1 := chrono.DurationFrom(time.Second * 5)
	d2 := chrono.DurationFrom(time.Second * 10)
	d3 := chrono.DurationFrom(time.Second * 5)

	assert.True(t, d1.SmallerOrEqual(d2), "d1 should be smaller or equal to d2")
	assert.True(t, d1.SmallerOrEqual(d3), "d1 should be smaller or equal to d3")
	assert.True(t, d2.SmallerOrEqual(d2), "d2 should be smaller or equal to itself")

	assert.True(t, d2.Greater(d1), "d2 should be greater than d1")
	assert.False(t, d1.Greater(d2), "d1 should not be greater than d2")
	assert.False(t, d1.Greater(d3), "d1 should not be greater than d3")

	assert.True(t, d2.GreaterOrEqual(d1), "d2 should be greater or equal to d1")
	assert.False(t, d1.GreaterOrEqual(d2), "d1 should not be greater or equal to d2")
	assert.True(t, d1.GreaterOrEqual(d3), "d1 should be greater or equal to d3")

	assert.True(t, d1.Smaller(d2), "d1 should be smaller than d2")
	assert.False(t, d2.Smaller(d1), "d2 should not be smaller than d1")
	assert.False(t, d2.Smaller(d2), "d2 should not be smaller than itself")
}

func TestStringConversion(t *testing.T) {
	d := chrono.DurationFrom(5 * time.Second)

	assert.Equal(t, "5s", d.String(), "String representation should match")
	assert.Equal(t, int64(5000000000), d.Nanoseconds(), "Nanoseconds should match")
	assert.Equal(t, int64(5000000), d.Microseconds(), "Microseconds should match")
	assert.Equal(t, int64(5000), d.Milliseconds(), "Milliseconds should match")
	assert.Equal(t, 5.0, d.Seconds(), "Seconds should match")
	assert.Equal(t, 5.0/60.0, d.Minutes(), "Minutes should match")
	assert.Equal(t, 5.0/3600.0, d.Hours(), "Hours should match")
}

func TestTruncateAndRound(t *testing.T) {
	d := chrono.DurationFrom(5*time.Second + 500*time.Millisecond)
	truncated := d.Truncate(chrono.DurationFrom(time.Second))
	rounded := d.Round(chrono.DurationFrom(time.Second))

	assert.Equal(t, chrono.DurationFrom(5*time.Second), truncated, "Truncate should work as expected")
	assert.Equal(t, chrono.DurationFrom(6*time.Second), rounded, "Round should work as expected")
}

func TestAbs(t *testing.T) {
	d1 := chrono.DurationFrom(-5 * time.Second)
	d2 := chrono.DurationFrom(5 * time.Second)

	assert.Equal(t, d2, d1.Abs(), "Absolute value should match")
}

func TestParseDuration(t *testing.T) {
	durationStr := "3h30m"
	expectedDuration := chrono.DurationFrom(3*time.Hour + 30*time.Minute)

	parsedDuration, err := chrono.ParseDuration(durationStr)
	assert.NoError(t, err, "Parsing should not return an error")
	assert.Equal(t, expectedDuration, parsedDuration, "Parsed duration should match")
}

func TestDynamoDBAttributeHandling(t *testing.T) {
	d := chrono.DurationFrom(5 * time.Second)

	attrValue, err := d.MarshalDynamoDBAttributeValue()
	assert.NoError(t, err, "MarshalDynamoDBAttributeValue should not return an error")

	var unmarshaledDuration chrono.Duration
	err = unmarshaledDuration.UnmarshalDynamoDBAttributeValue(attrValue)
	assert.NoError(t, err, "UnmarshalDynamoDBAttributeValue should not return an error")
	assert.Equal(t, d, unmarshaledDuration, "Unmarshaled duration should match")
}

func TestJSONMarshaling(t *testing.T) {
	d := chrono.DurationFrom(5 * time.Second)

	jsonData, err := d.MarshalJSON()
	assert.NoError(t, err, "MarshalJSON should not return an error")

	var unmarshaledDuration chrono.Duration
	err = unmarshaledDuration.UnmarshalJSON(jsonData)
	assert.NoError(t, err, "UnmarshalJSON should not return an error")
	assert.Equal(t, d, unmarshaledDuration, "Unmarshaled duration should match")
}
