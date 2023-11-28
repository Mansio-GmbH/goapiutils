package money

import (
	"encoding/json"

	"github.com/Rhymond/go-money"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type dynamodbMarshalMoneyWithoutVat struct {
	Amount       int64  `dynamodbav:"amount"`
	CurrencyCode string `dynamodbav:"currencyCode"`
	Display      string `dynamodbav:"display"`
}

func (m MoneyWithoutVat) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	mm := dynamodbMarshalMoneyWithoutVat{
		Amount:       m.Amount(),
		CurrencyCode: m.CurrencyCode(),
		Display:      m.Display(),
	}

	return attributevalue.Marshal(mm)
}

func (m *MoneyWithoutVat) UnmarshalDynamoDBAttributeValue(v types.AttributeValue) error {
	mm := dynamodbMarshalMoneyWithoutVat{}
	if err := attributevalue.Unmarshal(v, &mm); err != nil {
		return err
	}
	m.money = *money.New(mm.Amount, mm.CurrencyCode)
	return nil
}

func (m MoneyWithoutVat) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.money)
}

func (m *MoneyWithoutVat) UnmarshalJSON(b []byte) error {
	mon := money.Money{}
	if err := json.Unmarshal(b, &mon); err != nil {
		return err
	}
	m.money = mon
	return nil
}
