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
	require.Equal(t, `{"vatCode":"VAT_19_00","currencyCode":"EUR","amountNet":100.54,"amountGross":119.64,"leadingValueIsGross":false,"displayNet":"€100.54","displayGross":"€119.64"}`, string(b))
	m = money.NewFromGross(100_54, "EUR", money.VAT_19_00)
	b, err = json.Marshal(m)
	require.NoError(t, err)
	require.Equal(t, `{"vatCode":"VAT_19_00","currencyCode":"EUR","amountNet":84.49,"amountGross":100.54,"leadingValueIsGross":true,"displayNet":"€84.49","displayGross":"€100.54"}`, string(b))
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

/*
func TestMarshalFromAnyMap(t *testing.T) {
	m := money.NewFromNet(100_00, "EUR", money.VAT_19_00)
	mAsAnyMap, err := m.ToAnyMap()
	require.NoError(t, err)
	require.Equal(t, float64(100), mAsAnyMap["amountNet"])
	require.Equal(t, float64(119), mAsAnyMap["amountGross"])
	require.Equal(t, false, mAsAnyMap["leadingValueIsGross"])

	mFromAnyMap, err := money.MoneyFromAnyMap(mAsAnyMap)
	require.NoError(t, err)
	require.Equal(t, m.AmountNet(), mFromAnyMap.AmountNet())
	require.Equal(t, m.AmountGross(), mFromAnyMap.AmountGross())
}

func TestFromAny(t *testing.T) {
	m := money.NewFromNet(100_00, "EUR", money.VAT_19_00)
	mFromAny, tFromAny, err := money.FromAny(m)
	require.NoError(t, err)
	require.Nil(t, tFromAny)
	require.True(t, must.Must(mFromAny.Equals(m)))

	mFromAny, tFromAny, err = money.FromAny(*m)
	require.NoError(t, err)
	require.Nil(t, tFromAny)
	require.True(t, must.Must(mFromAny.Equals(m)))

	mAsAnyMap, err := m.ToAnyMap()
	require.NoError(t, err)
	mFromAny, tFromAny, err = money.FromAny(mAsAnyMap)
	require.NoError(t, err)
	require.Nil(t, tFromAny)
	require.True(t, must.Must(mFromAny.Equals(m)))

	total := money.NewTotal(m)
	require.NoError(t, err)
	mFromAny, tFromAny, err = money.FromAny(total)
	require.NoError(t, err)
	require.Nil(t, mFromAny)
	require.False(t, must.Must(total.GrossTotalOrZero("EUR").Equals(tFromAny.NetTotalOrZero("EUR"))))
	require.True(t, must.Must(total.NetTotalOrZero("EUR").Equals(tFromAny.NetTotalOrZero("EUR"))))

	require.NoError(t, err)
	mFromAny, tFromAny, err = money.FromAny(*total)
	require.NoError(t, err)
	require.Nil(t, mFromAny)
	require.False(t, must.Must(total.GrossTotalOrZero("EUR").Equals(tFromAny.NetTotalOrZero("EUR"))))
	require.True(t, must.Must(total.NetTotalOrZero("EUR").Equals(tFromAny.NetTotalOrZero("EUR"))))

	tAsAnyMap, err := total.ToAnyMap()
	require.NoError(t, err)
	mFromAny, tFromAny, err = money.FromAny(tAsAnyMap)
	require.NoError(t, err)
	require.Nil(t, mFromAny)
	require.False(t, must.Must(total.GrossTotalOrZero("EUR").Equals(tFromAny.NetTotalOrZero("EUR"))))
	require.True(t, must.Must(total.NetTotalOrZero("EUR").Equals(tFromAny.NetTotalOrZero("EUR"))))

	_, _, err = money.FromAny(12)
	require.Error(t, err)
	_, _, err = money.FromAny(make(map[string]any))
	require.Error(t, err)
}
*/
