package money_test

import (
	"testing"

	"github.com/mansio-gmbh/goapiutils/money"
	"github.com/stretchr/testify/require"
)

func TestSameVatRate(t *testing.T) {
	require.True(t, money.VAT_19_00.SameRate(money.VAT_19_00))
	require.False(t, money.VAT_19_00.SameRate(money.VAT_07_00))
}

func TestVatRateByCode(t *testing.T) {
	vat, err := money.VatByCode("VAT_19_00")
	require.NoError(t, err)
	require.True(t, vat.SameRate(money.VAT_19_00))

	vat, err = money.VatByCode("VAT_20_01")
	require.Error(t, err)
	require.Nil(t, vat)
}

func TestVatDisplay(t *testing.T) {
	require.Equal(t, "19%", money.VAT_19_00.Display())
	require.Equal(t, "7%", money.VAT_07_00.Display())
}
