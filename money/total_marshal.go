package money

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type jsonTotalMarshal struct {
	Nets        []*Money           `json:"nets" dynamodbav:"nets"`
	NetTotals   []*MoneyWithoutVat `json:"netTotals" dynamodbav:"netTotals"`
	GrossTotals []*MoneyWithoutVat `json:"grossTotals" dynamodbav:"grossTotals"`
	VatTotals   []*MoneyWithoutVat `json:"vatTotals" dynamodbav:"vatTotals"`
}

func (t Total) MarshalJSON() ([]byte, error) {
	nets := make([]*Money, 0)
	for _, vatByVatCode := range t.net {
		for vatCode, money := range vatByVatCode {
			vat, _ := VatByCode(vatCode)
			nets = append(nets, NewFromNet(money.Amount(), money.CurrencyCode(), vat))
		}
	}
	tm := jsonTotalMarshal{
		Nets:        nets,
		NetTotals:   t.NetTotals(),
		GrossTotals: t.GrossTotals(),
		VatTotals:   t.VatTotals(),
	}
	return json.Marshal(tm)
}

func (t *Total) UnmarshalJSON(b []byte) error {
	tm := jsonTotalMarshal{}
	if err := json.Unmarshal(b, &tm); err != nil {
		return err
	}
	t.net = make(map[string]map[vatcode]*MoneyWithoutVat)
	for _, net := range tm.Nets {
		if _, found := t.net[net.Currency()]; !found {
			t.net[net.Currency()] = make(map[vatcode]*MoneyWithoutVat)
		}
		t.net[net.Currency()][net.VAT()] = NewWithoutVat(net.Amount(), net.Currency())
	}
	return nil
}

type dynamodbTotalMarshal struct {
	Net map[string]map[string]int64 `dynamodbav:"net"`
}

func (t Total) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	tm := dynamodbTotalMarshal{
		Net: make(map[string]map[string]int64),
	}
	for currency, vatByVatCode := range t.net {
		if _, found := tm.Net[currency]; !found {
			tm.Net[currency] = make(map[string]int64)
		}
		for vatCode, money := range vatByVatCode {
			tm.Net[currency][vatCode] = money.Amount()
		}
	}

	return attributevalue.Marshal(tm)
}

func (t *Total) UnmarshalDynamoDBAttributeValue(v types.AttributeValue) error {
	tm := dynamodbTotalMarshal{}
	t.net = make(map[string]map[vatcode]*MoneyWithoutVat)
	if err := attributevalue.Unmarshal(v, &tm); err != nil {
		return err
	}
	for currency, amountByVatCode := range tm.Net {
		if _, found := t.net[currency]; !found {
			t.net[currency] = make(map[vatcode]*MoneyWithoutVat)
		}
		for vatCode, amount := range amountByVatCode {
			t.net[currency][vatCode] = NewWithoutVat(amount, currency)
		}
	}
	return nil
}
