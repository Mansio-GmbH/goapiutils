package ct

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/elliotchance/pie/v2"
	"github.com/mansio-gmbh/goapiutils/stringnormalisation"
)

type LicensePlate struct {
	licensePlateSegments []string
}

func (lp LicensePlate) MarshalJSON() ([]byte, error) {
	str := strings.Join(lp.licensePlateSegments, "-")
	return json.Marshal(str)
}

func (lp *LicensePlate) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	lp.licensePlateSegments = strings.FieldsFunc(str, func(r rune) bool {
		return r == '-' || r == ' '
	})
	return nil
}

func (lp *LicensePlate) UnmarshalDynamoDBAttributeValue(v types.AttributeValue) error {
	var str string
	if err := attributevalue.Unmarshal(v, &str); err != nil {
		return err
	}
	lp.licensePlateSegments = strings.FieldsFunc(str, func(r rune) bool {
		return r == '-' || r == ' '
	})
	return nil
}

func (lp LicensePlate) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return attributevalue.Marshal(lp.String())
}

func (lp LicensePlate) String() string {
	return strings.Join(pie.Map(lp.licensePlateSegments, func(segment string) string {
		return strings.ToUpper(stringnormalisation.NormaliseWithoutLengthCheck(segment))
	}), "-")
}

func NewLicensePlate(segments ...string) LicensePlate {
	return LicensePlate{licensePlateSegments: segments}
}

func ParseLicensePlate(s string) LicensePlate {
	return LicensePlate{licensePlateSegments: strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == ' '
	})}
}
