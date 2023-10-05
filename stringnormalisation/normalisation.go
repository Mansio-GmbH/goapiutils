package stringnormalisation

import (
	"errors"
	"regexp"
	"strings"
)

const (
	MIN_STRING_LENGTH = 5
)

var (
	ErrInputToShort  = errors.New("input too short")
	ErrOutputToShort = errors.New("output too short")
)

func Normalise(input string) (string, error) {
	if len(input) < MIN_STRING_LENGTH {
		return "", ErrInputToShort
	}
	output := NormaliseWithoutLengthCheck(input)
	if len(output) < MIN_STRING_LENGTH {
		return "", ErrOutputToShort
	}
	return output, nil
}

func NormaliseWithoutLengthCheck(input string) string {
	regex := regexp.MustCompile("[^a-zA-Z0-9-_~]")
	output := strings.ToUpper(regex.ReplaceAllString(input, ""))
	return output
}
