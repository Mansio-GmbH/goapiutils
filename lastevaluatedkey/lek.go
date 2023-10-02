package lastevaluatedkey

import (
	"encoding/base64"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Encode(lek map[string]types.AttributeValue) (*string, error) {
	if len(lek) == 0 {
		return nil, nil
	}
	lekVal := make(map[string]string)
	if err := attributevalue.UnmarshalMap(lek, &lekVal); err != nil {
		return nil, err
	}
	jsonBytes, err := json.Marshal(lekVal)
	if err != nil {
		return nil, err
	}
	s := base64.URLEncoding.EncodeToString(jsonBytes)
	return &s, nil
}

func Decode(lek *string) (map[string]types.AttributeValue, error) {
	if lek == nil || *lek == "" {
		return nil, nil
	}
	b, err := base64.URLEncoding.DecodeString(*lek)
	if err != nil {
		return nil, err
	}
	lekVal := make(map[string]string)
	if err := json.Unmarshal(b, &lekVal); err != nil {
		return nil, err
	}
	return attributevalue.MarshalMap(lekVal)
}
