package money

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type vatCodeAndMoney struct {
	Amount       float64 `json:"amount" dynamodbav:"amount"`
	CurrencyCode string  `json:"currencyCode" dynamodbav:"currencyCode"`
	Display      string  `json:"display" dynamodbav:"display"`
	VatCode      string  `json:"vatCode" dynamodbav:"vatCode"`
}

type totalMarshal struct {
	NetTotals    []*MoneyWithoutVat `json:"netTotals" dynamodbav:"netTotals"`
	GrossTotals  []*MoneyWithoutVat `json:"grossTotals" dynamodbav:"grossTotals"`
	VatTotals    []*MoneyWithoutVat `json:"vatTotals" dynamodbav:"vatTotals"`
	VatByVatCode []*vatCodeAndMoney `json:"vatByVatCode" dynamodbav:"vatByVatCode"`
}

func (t Total) toMarshal() totalMarshal {
	tm := totalMarshal{
		NetTotals:    t.NetTotals(),
		GrossTotals:  t.GrossTotals(),
		VatTotals:    t.VatTotals(),
		VatByVatCode: t.VatTotalsByCodes(),
	}
	return tm
}

func (t Total) VatTotalsByCodes() []*vatCodeAndMoney {
	r := make([]*vatCodeAndMoney, 0)
	for vatCode, moneys := range t.vatByCode {
		for _, money := range moneys {
			r = append(r, &vatCodeAndMoney{
				VatCode:      vatCode,
				Amount:       money.money.AsMajorUnits(),
				CurrencyCode: money.CurrencyCode(),
				Display:      money.Display(),
			})
		}
	}
	return r
}

func (t *Total) fromMarshal(tm totalMarshal) error {
	*t = *NewTotal()
	t.netTotal = make(map[string]*MoneyWithoutVat)
	for _, total := range tm.NetTotals {
		t.netTotal[total.CurrencyCode()] = total
	}
	t.grossTotal = make(map[string]*MoneyWithoutVat)
	for _, total := range tm.GrossTotals {
		t.grossTotal[total.CurrencyCode()] = total
	}
	t.vatTotal = make(map[string]*MoneyWithoutVat)
	for _, total := range tm.VatTotals {
		t.vatTotal[total.CurrencyCode()] = total
	}
	t.vatByCode = make(map[string]map[string]*MoneyWithoutVat)
	for _, vatByVatCode := range tm.VatByVatCode {
		if _, found := t.vatByCode[vatByVatCode.VatCode]; !found {
			t.vatByCode[vatByVatCode.VatCode] = make(map[string]*MoneyWithoutVat)
		}
		t.vatByCode[vatByVatCode.VatCode][vatByVatCode.CurrencyCode] = NewWithoutVatFromFloat(vatByVatCode.Amount, vatByVatCode.CurrencyCode)
	}
	return nil
}

func (t Total) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.toMarshal())
}

func (t *Total) UnmarshalJSON(b []byte) error {
	tm := totalMarshal{}
	if err := json.Unmarshal(b, &tm); err != nil {
		return err
	}
	return t.fromMarshal(tm)
}

func (t Total) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return attributevalue.Marshal(t.toMarshal())
}

func (t *Total) UnmarshalDynamoDBAttributeValue(v types.AttributeValue) error {
	tm := totalMarshal{}
	if err := attributevalue.Unmarshal(v, &tm); err != nil {
		return err
	}
	return t.fromMarshal(tm)
}

func (t *Total) ToAnyMap() (map[string]any, error) {
	asJson, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	val := make(map[string]any)
	err = json.Unmarshal(asJson, &val)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func TotalFromAnyMap(val map[string]any) (*Total, error) {
	hasNetTotals := false
	hasGrossTotals := false
	hasVatTotals := false
	hasVatByVatCodes := false

	_, hasNetTotals = val["netTotals"]
	_, hasGrossTotals = val["grossTotals"]
	_, hasVatTotals = val["vatTotals"]
	_, hasVatByVatCodes = val["vatByVatCode"]

	if !(hasNetTotals && hasGrossTotals && hasVatTotals && hasVatByVatCodes) {
		return nil, errors.New("value is no valid total")
	}

	asJson, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	m := Total{}
	err = json.Unmarshal(asJson, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
