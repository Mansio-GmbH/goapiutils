package money_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/mansio-gmbh/goapiutils/money"
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
	fmt.Println(jsonStr)
	require.Contains(t, jsonStr, `"netTotals":[`)
	require.Contains(t, jsonStr, `{"amount":600,"currencyCode":"EUR","display":"€600.00"}`)
	require.Contains(t, jsonStr, `{"amount":600,"currencyCode":"USD","display":"$600.00"}`)
	require.Contains(t, jsonStr, `"grossTotals":[`)
	require.Contains(t, jsonStr, `{"amount":652,"currencyCode":"EUR","display":"€652.00"}`)
	require.Contains(t, jsonStr, `{"amount":652,"currencyCode":"USD","display":"$652.00"}`)
	require.Contains(t, jsonStr, `"vatTotals":[`)
	require.Contains(t, jsonStr, `{"amount":52,"currencyCode":"EUR","display":"€52.00"}`)
	require.Contains(t, jsonStr, `{"amount":52,"currencyCode":"USD","display":"$52.00"}`)
}

func TestTotalUnmarshalJSON(t *testing.T) {
	jsonStr := `{"netTotals":[{"amount":600,"currencyCode":"EUR","display":"€600.00"},{"amount":600,"currencyCode":"USD","display":"$600.00"}],"grossTotals":[{"amount":652,"currencyCode":"EUR","display":"€652.00"},{"amount":652,"currencyCode":"USD","display":"$652.00"}],"vatTotals":[{"amount":52,"currencyCode":"EUR","display":"€52.00"},{"amount":52,"currencyCode":"USD","display":"$52.00"}]}`

	total := &money.Total{}
	err := json.Unmarshal([]byte(jsonStr), &total)
	require.Error(t, err)
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
	require.True(t, total.Equals(unmarshaledValue))
}
