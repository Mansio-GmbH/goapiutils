package money_test

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/mansio-gmbh/goapiutils/money"
	"github.com/mansio-gmbh/goapiutils/must"
	"github.com/stretchr/testify/require"
)

func TestMarshalJSON(t *testing.T) {
	m := money.NewFromNet(100_54, "EUR", money.VAT_19_00)
	b, err := json.Marshal(m)
	require.NoError(t, err)
	require.Equal(t, `{"vatCode":"VAT_19_00","currencyCode":"EUR","amountNet":100.54,"amountGross":119.64,"leadingValueIsGross":false,"displayNet":"€100.54","displayGross":"€119.64","amount":100.54,"valueIsGross":false}`, string(b))
	m = money.NewFromGross(100_54, "EUR", money.VAT_19_00)
	b, err = json.Marshal(m)
	require.NoError(t, err)
	require.Equal(t, `{"vatCode":"VAT_19_00","currencyCode":"EUR","amountNet":84.49,"amountGross":100.54,"leadingValueIsGross":true,"displayNet":"€84.49","displayGross":"€100.54","amount":100.54,"valueIsGross":true}`, string(b))
}

func TestUnmarshalJSON(t *testing.T) {
	data := `{"amount":100.00,"currencyCode":"EUR","vatCode":"VAT_19_00","valueIsGross":false}`
	m := money.Money{}
	err := json.Unmarshal([]byte(data), &m)
	require.NoError(t, err)
	require.Equal(t, m.AmountGross(), int64(119_00))
	require.Equal(t, m.AmountNet(), int64(100_00))
	require.Equal(t, m.Currency(), "EUR")
	require.Equal(t, m.VAT(), "VAT_19_00")
}

func TestMarshalDynamoDB(t *testing.T) {
	m := money.NewFromNet(100_00, "EUR", money.VAT_19_00)
	attr, err := attributevalue.Marshal(m)
	require.NoError(t, err)

	m2 := &money.Money{}
	attributevalue.Unmarshal(attr, m2)
	require.True(t, must.Must(m.Equals(m2)))
}

func TestUnmarshalJSONWithStrangeValues(t *testing.T) {
	data := `{"amount":4545.69,"currencyCode":"EUR","vatCode":"VAT_19_00","valueIsGross":false}`
	m := money.Money{}
	err := json.Unmarshal([]byte(data), &m)
	require.NoError(t, err)
	require.Equal(t, int64(4545_69), m.AmountNet())
	require.Equal(t, "EUR", m.Currency())
	require.Equal(t, "VAT_19_00", m.VAT())

	data = `{"amount":1128.60,"currencyCode":"EUR","vatCode":"VAT_19_00","valueIsGross":false}`
	err = json.Unmarshal([]byte(data), &m)
	require.NoError(t, err)
	require.Equal(t, int64(112860), m.AmountNet())
	require.Equal(t, "EUR", m.Currency())
	require.Equal(t, "VAT_19_00", m.VAT())
}
