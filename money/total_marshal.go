package money

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type jsonTotalMarshal struct {
	NetTotals   []*MoneyWithoutVat `json:"netTotals" dynamodbav:"netTotals"`
	GrossTotals []*MoneyWithoutVat `json:"grossTotals" dynamodbav:"grossTotals"`
	VatTotals   []*MoneyWithoutVat `json:"vatTotals" dynamodbav:"vatTotals"`
}

func (t Total) MarshalJSON() ([]byte, error) {
	tm := jsonTotalMarshal{
		NetTotals:   t.NetTotals(),
		GrossTotals: t.GrossTotals(),
		VatTotals:   t.VatTotals(),
	}
	return json.Marshal(tm)
}

func (t *Total) UnmarshalJSON(b []byte) error {
	return errors.New("not implemented")
}

type dynamodbTotalMarshal map[string]map[string]int64

func (t Total) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	tm := make(dynamodbTotalMarshal)
	for currency, vatByVatCode := range t.net {
		if _, found := tm[currency]; !found {
			tm[currency] = make(map[string]int64)
		}
		for vatCode, money := range vatByVatCode {
			tm[currency][vatCode] = money.Amount()
		}
	}

	return attributevalue.Marshal(tm)
}

func (t *Total) UnmarshalDynamoDBAttributeValue(v types.AttributeValue) error {
	tm := make(dynamodbTotalMarshal)
	t.net = make(map[string]map[vatcode]*MoneyWithoutVat)
	if err := attributevalue.Unmarshal(v, &tm); err != nil {
		return err
	}
	for currency, amountByVatCode := range tm {
		if _, found := t.net[currency]; !found {
			t.net[currency] = make(map[vatcode]*MoneyWithoutVat)
		}
		for vatCode, amount := range amountByVatCode {
			t.net[currency][vatCode] = NewWithoutVat(amount, currency)
		}
	}
	return nil
}
