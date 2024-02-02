package money_test

import (
	"testing"

	"github.com/mansio-gmbh/goapiutils/money"
	"github.com/mansio-gmbh/goapiutils/must"
	"github.com/stretchr/testify/require"
)

func TestTotalWithCurrencyAndVats(t *testing.T) {
	total := money.NewTotal()

	total.Add(money.NewFromGross(107_00, "EUR", money.VAT_07_00))
	total.Add(money.NewFromNet(100_00, "EUR", money.VAT_07_00))
	total.Add(money.NewFromGross(119_00, "EUR", money.VAT_19_00))
	total.Add(money.NewFromNet(100_00, "EUR", money.VAT_19_00))
	total.Add(money.NewFromGross(100_00, "EUR", money.VAT_00_00))
	total.Add(money.NewFromNet(100_00, "EUR", money.VAT_00_00))
	total.Add(money.NewFromGross(107_00, "USD", money.VAT_07_00))
	total.Add(money.NewFromNet(100_00, "USD", money.VAT_07_00))
	total.Add(money.NewFromGross(119_00, "USD", money.VAT_19_00))
	total.Add(money.NewFromNet(100_00, "USD", money.VAT_19_00))
	total.Add(money.NewFromGross(100_00, "USD", money.VAT_00_00))
	total.Add(money.NewFromNet(100_00, "USD", money.VAT_00_00))

	netSum, found := total.NetTotal("EUR")
	require.True(t, found)
	require.Equal(t, int64(600_00), netSum.Amount())

	netSum, found = total.NetTotal("USD")
	require.True(t, found)
	require.Equal(t, int64(600_00), netSum.Amount())

	netSum, found = total.NetTotal("GBP")
	require.False(t, found)
	require.Nil(t, netSum)

	grossSum, found := total.GrossTotal("EUR")
	require.True(t, found)
	require.Equal(t, int64(652_00), grossSum.Amount())

	grossSum, found = total.GrossTotal("USD")
	require.True(t, found)
	require.Equal(t, int64(652_00), grossSum.Amount())

	grossSum, found = total.GrossTotal("GBP")
	require.False(t, found)
	require.Nil(t, grossSum)

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

func TestTotalAddTotal(t *testing.T) {
	total1 := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	total2 := money.NewTotal(
		money.NewFromNet(100_00, "EUR", money.VAT_07_00),
		money.NewFromNet(100_00, "GBP", money.VAT_19_00),
		money.NewFromNet(20000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_00_00),
	)

	sumTotal := money.NewTotal2(total1, total2)
	grossTotal, found := sumTotal.GrossTotal("EUR")
	require.True(t, found)
	require.True(t, must.Must(grossTotal.Equals(money.NewWithoutVat(214_00, "EUR"))))
	netTotal, found := sumTotal.NetTotal("EUR")
	require.True(t, found)
	require.True(t, must.Must(netTotal.Equals(money.NewWithoutVat(200_00, "EUR"))))
	vat_07_00, found := sumTotal.VatTotalByCode("EUR", "VAT_07_00")
	require.True(t, found)
	require.True(t, must.Must(vat_07_00.Equals(money.NewWithoutVat(14_00, "EUR"))))
	vat_19_00, found := sumTotal.VatTotalByCode("EUR", "VAT_19_00")
	require.False(t, found)
	require.Nil(t, vat_19_00)
	vatTotal, found := sumTotal.VatTotal("EUR")
	require.True(t, found)
	require.True(t, must.Must(vatTotal.Equals(money.NewWithoutVat(14_00, "EUR"))))

	grossTotal, found = sumTotal.GrossTotal("USD")
	require.True(t, found)
	require.True(t, must.Must(grossTotal.Equals(money.NewWithoutVat(100_00, "USD"))))
	netTotal, found = sumTotal.NetTotal("USD")
	require.True(t, found)
	require.True(t, must.Must(netTotal.Equals(money.NewWithoutVat(100_00, "USD"))))

	grossTotal, found = sumTotal.GrossTotal("GBP")
	require.True(t, found)
	require.True(t, must.Must(grossTotal.Equals(money.NewWithoutVat(119_00, "GBP"))))
	netTotal, found = sumTotal.NetTotal("GBP")
	require.True(t, found)
	require.True(t, must.Must(netTotal.Equals(money.NewWithoutVat(100_00, "GBP"))))

	grossTotal, found = sumTotal.GrossTotal("JPY")
	require.True(t, found)
	require.True(t, must.Must(grossTotal.Equals(money.NewWithoutVat(56400_00, "JPY"))))
	netTotal, found = sumTotal.NetTotal("JPY")
	require.True(t, found)
	require.True(t, must.Must(netTotal.Equals(money.NewWithoutVat(50000_00, "JPY"))))
	vat_07_00, found = sumTotal.VatTotalByCode("JPY", "VAT_07_00")
	require.True(t, found)
	require.True(t, must.Must(vat_07_00.Equals(money.NewWithoutVat(700_00, "JPY"))))
	vat_19_00, found = sumTotal.VatTotalByCode("JPY", "VAT_19_00")
	require.True(t, found)
	require.True(t, must.Must(vat_19_00.Equals(money.NewWithoutVat(5700_00, "JPY"))))
	vat_00_00, found := sumTotal.VatTotalByCode("JPY", "VAT_00_00")
	require.True(t, found)
	require.True(t, must.Must(vat_00_00.Equals(money.NewWithoutVat(0_00, "JPY"))))
}

func TestTotalAddIndividually(t *testing.T) {
	total1 := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	total2 := total1.AddIndividually(money.NewWithoutVat(200_00, "EUR"), money.NewWithoutVat(226_00, "EUR"), money.NewWithoutVat(38_00, "EUR"), map[string]*money.MoneyWithoutVat{
		"VAT_19_00": money.NewWithoutVat(19, "EUR"),
		"VAT_07_00": money.NewWithoutVat(7, "EUR"),
	})

	netTotal, found := total2.NetTotal("EUR")
	require.True(t, found)
	require.True(t, must.Must(netTotal.Equals(money.NewWithoutVat(300_00, "EUR"))))

	grossTotal, found := total2.GrossTotal("EUR")
	require.True(t, found)
	require.True(t, must.Must(grossTotal.Equals(money.NewWithoutVat(333_00, "EUR"))))
}

func TestTotalAddIndividually_NewCurrency(t *testing.T) {
	total1 := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	total2 := total1.AddIndividually(money.NewWithoutVat(200_00, "GBP"), money.NewWithoutVat(226_00, "GBP"), money.NewWithoutVat(38_00, "GBP"), map[string]*money.MoneyWithoutVat{
		"VAT_19_00": money.NewWithoutVat(19, "GBP"),
		"VAT_07_00": money.NewWithoutVat(7, "GBP"),
	})

	netTotal, found := total2.NetTotal("GBP")
	require.True(t, found)
	require.True(t, must.Must(netTotal.Equals(money.NewWithoutVat(200_00, "GBP"))))

	grossTotal, found := total2.GrossTotal("GBP")
	require.True(t, found)
	require.True(t, must.Must(grossTotal.Equals(money.NewWithoutVat(226_00, "GBP"))))

	netTotal, found = total2.NetTotal("EUR")
	require.True(t, found)
	require.True(t, must.Must(netTotal.Equals(money.NewWithoutVat(100_00, "EUR"))))

	grossTotal, found = total2.GrossTotal("EUR")
	require.True(t, found)
	require.True(t, must.Must(grossTotal.Equals(money.NewWithoutVat(107_00, "EUR"))))
}

func TestXOrZero(t *testing.T) {
	total := money.NewTotal2()

	require.True(t, total.NetTotalOrZero("EUR").IsZero())
	require.True(t, total.GrossTotalOrZero("EUR").IsZero())
	require.True(t, total.VatTotalOrZero("EUR").IsZero())
	require.True(t, total.VatTotalByCodeOrZero("EUR", "VAT_00_00").IsZero())
}

func TestTotalNegate(t *testing.T) {
	total := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	total2 := total.Negate()
	require.Equal(t, int64(-107_00), total2.GrossTotalOrZero("EUR").Amount())
	require.Equal(t, int64(-100_00), total2.GrossTotalOrZero("USD").Amount())
	require.Equal(t, int64(-22600_00), total2.GrossTotalOrZero("JPY").Amount())
}

func TestTotalNegative(t *testing.T) {
	total := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(-10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(-10000_00, "JPY", money.VAT_07_00),
	)

	total2 := total.Negative()
	require.Equal(t, int64(-107_00), total2.GrossTotalOrZero("EUR").Amount())
	require.Equal(t, int64(-100_00), total2.GrossTotalOrZero("USD").Amount())
	require.Equal(t, int64(-22600_00), total2.GrossTotalOrZero("JPY").Amount())
}

func TestTotalAbsolute(t *testing.T) {
	total := money.NewTotal(
		money.NewFromGross(-107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(-10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(-10000_00, "JPY", money.VAT_07_00),
	)

	total2 := total.Absolute()
	require.Equal(t, int64(107_00), total2.GrossTotalOrZero("EUR").Amount())
	require.Equal(t, int64(100_00), total2.GrossTotalOrZero("USD").Amount())
	require.Equal(t, int64(22600_00), total2.GrossTotalOrZero("JPY").Amount())
}

func TestTotalForEachCurrency(t *testing.T) {
	total := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	currencies := make(map[string]bool)
	total.ForEachCurrency(func(currencyCode string, _, _, _ *money.MoneyWithoutVat, _ map[string]*money.MoneyWithoutVat) {
		currencies[currencyCode] = true
	})

	require.Equal(t, 3, len(currencies))
	require.True(t, currencies["EUR"])
	require.True(t, currencies["USD"])
	require.True(t, currencies["JPY"])
}

func TestTotalCurrencyCodes(t *testing.T) {
	total := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	currencyCodes := total.CurrencyCodes()
	require.Equal(t, 3, len(currencyCodes))
	require.Contains(t, currencyCodes, "EUR")
	require.Contains(t, currencyCodes, "USD")
	require.Contains(t, currencyCodes, "JPY")
}

func TestTotalVatCodes(t *testing.T) {
	total := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	vatCodes := total.VatCodes()

	require.Equal(t, 3, len(vatCodes))
	require.Contains(t, vatCodes, "VAT_07_00")
	require.Contains(t, vatCodes, "VAT_19_00")
	require.Contains(t, vatCodes, "VAT_00_00")
}

func TestTotalMultiply(t *testing.T) {
	total := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	total2 := total.Multiply(2)
	require.Equal(t, int64(214_00), total2.GrossTotalOrZero("EUR").Amount())
	require.Equal(t, int64(200_00), total2.GrossTotalOrZero("USD").Amount())
	require.Equal(t, int64(45200_00), total2.GrossTotalOrZero("JPY").Amount())
}

func TestTotalMultiplyByFloat(t *testing.T) {
	total := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	total2 := total.MultiplyByFloat(2.5)
	require.Equal(t, int64(267_50), total2.GrossTotalOrZero("EUR").Amount())
	require.Equal(t, int64(250_00), total2.GrossTotalOrZero("USD").Amount())
	require.Equal(t, int64(56500_00), total2.GrossTotalOrZero("JPY").Amount())
}

func TestTotalPercentage(t *testing.T) {
	total := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	total2 := total.Percentage(50)
	require.Equal(t, int64(53_50), total2.GrossTotalOrZero("EUR").Amount())
	require.Equal(t, int64(50_00), total2.GrossTotalOrZero("USD").Amount())
	require.Equal(t, int64(11300_00), total2.GrossTotalOrZero("JPY").Amount())
}

func TestTotalEquals(t *testing.T) {
	total1 := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	total2 := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	require.True(t, total1.Equals(total2))
}

func TestTotalEquals_false(t *testing.T) {
	total1 := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_07_00),
	)

	total2 := money.NewTotal(
		money.NewFromGross(107_00, "EUR", money.VAT_07_00),
		money.NewFromGross(100_00, "USD", money.VAT_00_00),
		money.NewFromNet(10000_00, "JPY", money.VAT_19_00),
	)

	require.False(t, total1.Equals(total2))
}
