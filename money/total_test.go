package money_test

import (
	"testing"

	"github.com/mansio-gmbh/api/lib/money"
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
