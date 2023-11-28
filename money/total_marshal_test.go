package money_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/mansio-gmbh/api/lib/money"
	"github.com/stretchr/testify/require"
)

func TestTotalMarshalJSON(t *testing.T) {
	total := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromNet(100_00, "EUR", money.VAT_07_00),
		money.NewFromGross(119_00, "EUR", money.VAT_19_00),
		money.NewFromNet(100_00, "EUR", money.VAT_19_00),
		money.NewFromGross(100_00, "EUR", money.VAT_00_00),
		money.NewFromNet(100_00, "EUR", money.VAT_00_00),
		money.NewFromGross(107_00, "USD", money.VAT_07_00),
		money.NewFromNet(100_00, "USD", money.VAT_07_00),
		money.NewFromGross(119_00, "USD", money.VAT_19_00),
		money.NewFromNet(100_00, "USD", money.VAT_19_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(100_00, "USD", money.VAT_00_00),
	)

	b, err := json.Marshal(total)
	require.NoError(t, err)
	jsonStr := string(b)
	require.Contains(t, jsonStr, `"netTotals":[`)
	require.Contains(t, jsonStr, `{"amount":600,"currencyCode":"EUR","display":"€600.00"}`)
	require.Contains(t, jsonStr, `{"amount":600,"currencyCode":"USD","display":"$600.00"}`)
	require.Contains(t, jsonStr, `"grossTotals":[`)
	require.Contains(t, jsonStr, `{"amount":652,"currencyCode":"EUR","display":"€652.00"}`)
	require.Contains(t, jsonStr, `{"amount":652,"currencyCode":"USD","display":"$652.00"}`)
	require.Contains(t, jsonStr, `"vatTotals":[`)
	require.Contains(t, jsonStr, `{"amount":52,"currencyCode":"EUR","display":"€52.00"}`)
	require.Contains(t, jsonStr, `{"amount":52,"currencyCode":"USD","display":"$52.00"}`)
	require.Contains(t, jsonStr, `"vatByVatCode":[`)
	require.Contains(t, jsonStr, `{"amount":14,"currencyCode":"EUR","display":"€14.00","vatCode":"VAT_07_00"}`)
	require.Contains(t, jsonStr, `{"amount":38,"currencyCode":"EUR","display":"€38.00","vatCode":"VAT_19_00"}`)
	require.Contains(t, jsonStr, `{"amount":14,"currencyCode":"USD","display":"$14.00","vatCode":"VAT_07_00"}`)
	require.Contains(t, jsonStr, `{"amount":38,"currencyCode":"USD","display":"$38.00","vatCode":"VAT_19_00"}`)

	fmt.Println(jsonStr)
}

func TestTotalUnmarshalJSON(t *testing.T) {
	jsonStr := `{"netTotals":[{"amount":600,"currencyCode":"EUR","display":"€600.00"},{"amount":600,"currencyCode":"USD","display":"$600.00"}],"grossTotals":[{"amount":652,"currencyCode":"EUR","display":"€652.00"},{"amount":652,"currencyCode":"USD","display":"$652.00"}],"vatTotals":[{"amount":52,"currencyCode":"EUR","display":"€52.00"},{"amount":52,"currencyCode":"USD","display":"$52.00"}],"vatByVatCode":[{"amount":14,"currencyCode":"EUR","display":"€14.00","vatCode":"VAT_07_00"},{"amount":14,"currencyCode":"USD","display":"$14.00","vatCode":"VAT_07_00"},{"amount":38,"currencyCode":"EUR","display":"€38.00","vatCode":"VAT_19_00"},{"amount":38,"currencyCode":"USD","display":"$38.00","vatCode":"VAT_19_00"},{"amount":0,"currencyCode":"EUR","display":"€0.00","vatCode":"VAT_00_00"},{"amount":0,"currencyCode":"USD","display":"$0.00","vatCode":"VAT_00_00"}]}`

	total := &money.Total{}
	err := json.Unmarshal([]byte(jsonStr), &total)
	require.NoError(t, err)

	netSum, found := total.NetTotal("EUR")
	require.True(t, found)
	require.Equal(t, int64(600_00), netSum.Amount())

	netSum, found = total.NetTotal("USD")
	require.True(t, found)
	require.Equal(t, int64(600_00), netSum.Amount())

	netSum, found = total.NetTotal("GBP")
	require.False(t, found)
	require.Nil(t, netSum)

	vatIncludedTotal, found := total.VatTotal("EUR")
	require.True(t, found)
	require.Equal(t, int64(52_00), vatIncludedTotal.Amount())

	vatIncludedTotal, found = total.VatTotal("USD")
	require.True(t, found)
	require.Equal(t, int64(52_00), vatIncludedTotal.Amount())

	vatIncluded, found := total.VatTotalByCode("EUR", "VAT_07_00")
	require.True(t, found)
	require.Equal(t, int64(14_00), vatIncluded.Amount())

	vatIncluded, found = total.VatTotalByCode("USD", "VAT_07_00")
	require.True(t, found)
	require.Equal(t, int64(14_00), vatIncluded.Amount())

	vatIncluded, found = total.VatTotalByCode("EUR", "VAT_19_00")
	require.True(t, found)
	require.Equal(t, int64(38_00), vatIncluded.Amount())

	vatIncluded, found = total.VatTotalByCode("USD", "VAT_19_00")
	require.True(t, found)
	require.Equal(t, int64(38_00), vatIncluded.Amount())

	vatIncluded, found = total.VatTotalByCode("EUR", "VAT_00_00")
	require.True(t, found)
	require.Equal(t, int64(0), vatIncluded.Amount())

	vatIncluded, found = total.VatTotalByCode("USD", "VAT_00_00")
	require.True(t, found)
	require.Equal(t, int64(0), vatIncluded.Amount())
}

func TestTotalMarshalDynamoDB(t *testing.T) {
	total := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromNet(100_00, "EUR", money.VAT_07_00),
		money.NewFromGross(119_00, "EUR", money.VAT_19_00),
		money.NewFromNet(100_00, "EUR", money.VAT_19_00),
		money.NewFromGross(100_00, "EUR", money.VAT_00_00),
		money.NewFromNet(100_00, "EUR", money.VAT_00_00),
		money.NewFromGross(107_00, "USD", money.VAT_07_00),
		money.NewFromNet(100_00, "USD", money.VAT_07_00),
		money.NewFromGross(119_00, "USD", money.VAT_19_00),
		money.NewFromNet(100_00, "USD", money.VAT_19_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(100_00, "USD", money.VAT_00_00),
	)

	marshalledValues, err := attributevalue.Marshal(total)
	require.NoError(t, err)

	unmarshaledValue := &money.Total{}
	err = attributevalue.Unmarshal(marshalledValues, &unmarshaledValue)
	require.NoError(t, err)
}

func TestEmptyMarshal(t *testing.T) {
	total := money.NewTotal()

	marshalledValues, err := json.MarshalIndent(total, "", "  ")
	require.NoError(t, err)

	fmt.Println(string(marshalledValues))

	unmarshaledValue := &money.Total{}
	err = json.Unmarshal(marshalledValues, &unmarshaledValue)
	require.NoError(t, err)
}
