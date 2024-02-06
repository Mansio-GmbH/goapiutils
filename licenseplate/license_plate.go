package licenseplate

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type LicensePlate struct {
	LicensePlate []string `json:"licensePlate"`
}

func (lp LicensePlate) MarshalJSON() ([]byte, error) {
	str := strings.Join(lp.LicensePlate, "-")
	return json.Marshal(str)
}

func (lp *LicensePlate) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	lp.LicensePlate = strings.FieldsFunc(str, func(r rune) bool {
		return r == '-' || r == ' '
	})
	return nil
}

func (lp *LicensePlate) UnmarshalDynamoDBAttributeValue(v types.AttributeValue) error {
	var str string
	if err := attributevalue.Unmarshal(v, &str); err != nil {
		return err
	}
	lp.LicensePlate = strings.FieldsFunc(str, func(r rune) bool {
		return r == '-' || r == ' '
	})
	return nil
}

func (lp LicensePlate) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return attributevalue.Marshal(lp.LicensePlate)
}
