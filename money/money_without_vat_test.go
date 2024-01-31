package money_test

import (
	"testing"

	gomoney "github.com/Rhymond/go-money"
	"github.com/mansio-gmbh/goapiutils/money"
	"github.com/stretchr/testify/require"
)

func TestNewMoneyWithoutVat(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	require.Equal(t, int64(100), m.Amount())
	require.Equal(t, "EUR", m.CurrencyCode())
}

func TestNewMoneyWithoutVatFromFloat(t *testing.T) {
	m := money.NewWithoutVatFromFloat(100.0, "EUR")
	require.Equal(t, int64(10000), m.Amount())
	require.Equal(t, "EUR", m.CurrencyCode())
}

func TestNewMoneyWithoutVatFromMoney(t *testing.T) {
	m := money.NewWithoutVatFromMoney(gomoney.New(100, "EUR"))
	require.Equal(t, int64(100), m.Amount())
	require.Equal(t, "EUR", m.CurrencyCode())
}

func TestMoneyWithoutVat_Add(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	om := money.NewWithoutVat(200, "EUR")
	rm, err := m.Add(om)
	require.NoError(t, err)
	require.Equal(t, int64(300), rm.Amount())
	require.Equal(t, "EUR", rm.CurrencyCode())
}

func TestMoneyWithoutVat_Add_nil(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	rm, err := m.Add(nil)
	require.NoError(t, err)
	require.Equal(t, int64(100), rm.Amount())
	require.Equal(t, "EUR", rm.CurrencyCode())
}

func TestMoneyWithoutVat_Multiply(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	rm := m.Multiply(2)
	require.Equal(t, int64(200), rm.Amount())
	require.Equal(t, "EUR", rm.CurrencyCode())
}

func TestMoneyWithoutVat_MultiplyByFloat(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	rm := m.MultiplyByFloat(2.5)
	require.Equal(t, int64(250), rm.Amount())
	require.Equal(t, "EUR", rm.CurrencyCode())
}

func TestMoneyWithoutVat_Percentage(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	rm := m.Percentage(50)
	require.Equal(t, int64(50), rm.Amount())
	require.Equal(t, "EUR", rm.CurrencyCode())
}

func TestMoneyWithoutVat_Equals(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	om := money.NewWithoutVat(100, "EUR")
	eq, err := m.Equals(om)
	require.NoError(t, err)
	require.True(t, eq)
}

func TestMoneyWithoutVat_Equals_false(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	om := money.NewWithoutVat(200, "EUR")
	eq, err := m.Equals(om)
	require.NoError(t, err)
	require.False(t, eq)
}

func TestMoneyWithoutVat_IsZero(t *testing.T) {
	m := money.NewWithoutVat(0, "EUR")
	require.True(t, m.IsZero())
}

func TestMoneyWithoutVat_IsZero_false(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	require.False(t, m.IsZero())
}

func TestMoneyWithoutVat_IsPositive(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	require.True(t, m.IsPositive())
}

func TestMoneyWithoutVat_IsPositive_false(t *testing.T) {
	m := money.NewWithoutVat(-100, "EUR")
	require.False(t, m.IsPositive())
}

func TestMoneyWithoutVat_IsNegative(t *testing.T) {
	m := money.NewWithoutVat(-100, "EUR")
	require.True(t, m.IsNegative())
}

func TestMoneyWithoutVat_IsNegative_false(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	require.False(t, m.IsNegative())
}

func TestMoneyWithoutVat_LessThan(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	om := money.NewWithoutVat(200, "EUR")
	lt, err := m.LessThan(om)
	require.NoError(t, err)
	require.True(t, lt)
}

func TestMoneyWithoutVat_LessThan_false(t *testing.T) {
	m := money.NewWithoutVat(200, "EUR")
	om := money.NewWithoutVat(100, "EUR")
	lt, err := m.LessThan(om)
	require.NoError(t, err)
	require.False(t, lt)
}

func TestMoneyWithoutVat_LessThanOrEqual(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	om := money.NewWithoutVat(200, "EUR")
	lt, err := m.LessThanOrEqual(om)
	require.NoError(t, err)
	require.True(t, lt)
}

func TestMoneyWithoutVat_LessThanOrEqual_equal(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	om := money.NewWithoutVat(100, "EUR")
	lt, err := m.LessThanOrEqual(om)
	require.NoError(t, err)
	require.True(t, lt)
}

func TestMoneyWithoutVat_LessThanOrEqual_false(t *testing.T) {
	m := money.NewWithoutVat(200, "EUR")
	om := money.NewWithoutVat(100, "EUR")
	lt, err := m.LessThanOrEqual(om)
	require.NoError(t, err)
	require.False(t, lt)
}

func TestMoneyWithoutVat_GreaterThan(t *testing.T) {
	m := money.NewWithoutVat(200, "EUR")
	om := money.NewWithoutVat(100, "EUR")
	gt, err := m.GreaterThan(om)
	require.NoError(t, err)
	require.True(t, gt)
}

func TestMoneyWithoutVat_GreaterThan_false(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	om := money.NewWithoutVat(200, "EUR")
	gt, err := m.GreaterThan(om)
	require.NoError(t, err)
	require.False(t, gt)
}

func TestMoneyWithoutVat_GreaterThanOrEqual(t *testing.T) {
	m := money.NewWithoutVat(200, "EUR")
	om := money.NewWithoutVat(100, "EUR")
	gt, err := m.GreaterThanOrEqual(om)
	require.NoError(t, err)
	require.True(t, gt)
}

func TestMoneyWithoutVat_GreaterThanOrEqual_equal(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	om := money.NewWithoutVat(100, "EUR")
	gt, err := m.GreaterThanOrEqual(om)
	require.NoError(t, err)
	require.True(t, gt)
}

func TestMoneyWithoutVat_GreaterThanOrEqual_false(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	om := money.NewWithoutVat(200, "EUR")
	gt, err := m.GreaterThanOrEqual(om)
	require.NoError(t, err)
	require.False(t, gt)
}

func TestMoneyNegate(t *testing.T) {
	m := money.NewWithoutVat(100, "EUR")
	nm := m.Negate()
	require.Equal(t, int64(-100), nm.Amount())
	require.Equal(t, "EUR", nm.CurrencyCode())
}
