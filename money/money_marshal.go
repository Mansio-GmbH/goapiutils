package money

import (
	"encoding/json"

	"github.com/Rhymond/go-money"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type marshalVat struct {
	Rate        int64  `json:"rate" dynamodbav:"rate"`
	Denominator int64  `json:"denominator" dynamodbav:"denominator"`
	Code        string `json:"code" dynamodbav:"code"`
}

type dynamodbMarshalMoney struct {
	AmountGross         int64       `dynamodbav:"amountGrossCents"`
	AmountNet           int64       `dynamodbav:"amountNetCents"`
	CurrenyCode         string      `dynamodbav:"currencyCode"`
	Vat                 *marshalVat `dynamodbav:"vat"`
	LeadingValueIsGross bool        `dynamodbav:"leadingValueIsGross"`

	// output
	AmountGrossMajorUnits float64 `dynamodbav:"amountGross"`
	AmountNetMajorUnits   float64 `dynamodbav:"amountNet"`
	VatCode               string  `dynamodbav:"vatCode"`
	DisplayNet            string  `dynamodbav:"displayNet"`
	DisplayGross          string  `dynamodbav:"displayGross"`
}

type jsonMarshalMoney struct {
	VatCode     string `json:"vatCode"`
	CurrenyCode string `json:"currencyCode"`
	// output only
	AmountNet           float64 `json:"amountNet"`
	AmountGross         float64 `json:"amountGross"`
	LeadingValueIsGross bool    `json:"leadingValueIsGross"`
	DisplayNet          string  `json:"displayNet"`
	DisplayGross        string  `json:"displayGross"`
	// input only
	Amount       *float64 `json:"amount,omitempty" validate:"required"`
	ValueIsGross *bool    `json:"valueIsGross,omitempty" validate:"required"`
}

type jsonMarshalMoneyWithoutVat struct {
	Amount      float64 `json:"amount"`
	CurrenyCode string  `json:"currencyCode"`
	Display     string  `json:"display"`
}

func (m Money) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	mm := dynamodbMarshalMoney{
		AmountGross:           m.grossMoney.Amount(),
		AmountNet:             m.netMoney.Amount(),
		AmountGrossMajorUnits: m.grossMoney.AsMajorUnits(),
		AmountNetMajorUnits:   m.netMoney.AsMajorUnits(),
		CurrenyCode:           m.netMoney.Currency().Code,
		LeadingValueIsGross:   m.LeadingValueIsGross,
		VatCode:               m.vat.code,
		DisplayNet:            m.netMoney.Display(),
		DisplayGross:          m.grossMoney.Display(),
		Vat: &marshalVat{
			Rate:        m.vat.rate,
			Denominator: m.vat.denominator,
			Code:        m.vat.code,
		},
	}
	return attributevalue.Marshal(mm)
}

func (m *Money) UnmarshalDynamoDBAttributeValue(v types.AttributeValue) error {
	mm := dynamodbMarshalMoney{}
	if err := attributevalue.Unmarshal(v, &mm); err != nil {
		return err
	}
	m.grossMoney = money.New(mm.AmountGross, mm.CurrenyCode)
	m.netMoney = money.New(mm.AmountNet, mm.CurrenyCode)
	m.LeadingValueIsGross = mm.LeadingValueIsGross
	m.vat = &vat{
		rate:        mm.Vat.Rate,
		denominator: mm.Vat.Denominator,
		code:        mm.Vat.Code,
	}
	return nil
}

func (m Money) MarshalJSON() ([]byte, error) {
	leadingAmount := m.AmountNetAsMajorUnits()
	if m.LeadingValueIsGross {
		leadingAmount = m.AmountGrossAsMajorUnits()
	}
	mm := jsonMarshalMoney{
		Amount:              &leadingAmount,
		ValueIsGross:        &m.LeadingValueIsGross,
		AmountNet:           m.AmountNetAsMajorUnits(),
		AmountGross:         m.AmountGrossAsMajorUnits(),
		CurrenyCode:         m.Currency(),
		VatCode:             m.vat.code,
		LeadingValueIsGross: m.LeadingValueIsGross,
		DisplayNet:          m.DisplayNet(),
		DisplayGross:        m.DisplayGross(),
	}
	return json.Marshal(mm)
}

func (m *Money) UnmarshalJSON(b []byte) error {
	mm := jsonMarshalMoney{}
	if err := json.Unmarshal(b, &mm); err != nil {
		return err
	}

	vat, err := VatByCode(mm.VatCode)
	if err != nil {
		return err
	}
	var mc *Money
	if mm.Amount != nil && mm.ValueIsGross != nil {
		mc = NewFromFloat(*mm.Amount, mm.CurrenyCode, vat, *mm.ValueIsGross)
	} else if mm.LeadingValueIsGross {
		mc = NewFromFloat(mm.AmountGross, mm.CurrenyCode, vat, true)
	} else {
		mc = NewFromFloat(mm.AmountNet, mm.CurrenyCode, vat, false)
	}

	m.grossMoney = mc.grossMoney
	m.netMoney = mc.netMoney
	m.vat = mc.vat
	m.LeadingValueIsGross = mc.LeadingValueIsGross

	return nil
}

func init() {
	money.MarshalJSON = func(m money.Money) ([]byte, error) {
		mm := jsonMarshalMoneyWithoutVat{
			Amount:      m.AsMajorUnits(),
			CurrenyCode: m.Currency().Code,
			Display:     m.Display(),
		}
		return json.Marshal(mm)
	}
	money.UnmarshalJSON = func(m *money.Money, b []byte) error {
		mm := jsonMarshalMoneyWithoutVat{}

		if err := json.Unmarshal(b, &mm); err != nil {
			return err
		}
		*m = *money.NewFromFloat(mm.Amount, mm.CurrenyCode)
		return nil
	}
}
