package money_test

import (
	"fmt"
	"testing"

	"github.com/mansio-gmbh/api/lib/money"
	"github.com/mansio-gmbh/goapiutils/must"
	"github.com/stretchr/testify/require"
)

func TestNewMoney(t *testing.T) {
	m := money.NewFromNet(100_00, "EUR", money.VAT_19_00)
	require.Equal(t, int64(119_00), m.AmountGross(), "19: Gross from net")
	require.Equal(t, int64(100_00), m.AmountNet(), "19: Net from net")

	m = money.NewFromGross(100_00, "EUR", money.VAT_19_00)
	require.Equal(t, int64(100_00), m.AmountGross(), "19: Gross from gross")
	require.Equal(t, int64(84_03), m.AmountNet(), "19: Net from gross")

	m = money.NewFromNet(100_00, "EUR", money.VAT_07_00)

	require.Equal(t, int64(107_00), m.AmountGross(), "7: Gross from net")
	require.Equal(t, int64(100_00), m.AmountNet(), "7: Net from net")

	m = money.NewFromGross(100_00, "EUR", money.VAT_07_00)
	require.Equal(t, int64(100_00), m.AmountGross(), "7: Gross from gross")
	require.Equal(t, int64(93_46), m.AmountNet(), "7: Net from gross")

	m = money.New(100_00, "EUR", money.VAT_19_00, true)
	require.Equal(t, int64(100_00), m.AmountGross(), "19: Gross from gross")
	require.Equal(t, int64(84_03), m.AmountNet(), "19: Net from gross")

	m = money.New(100_00, "EUR", money.VAT_19_00, false)
	require.Equal(t, int64(119_00), m.AmountGross(), "19: Gross from net")
	require.Equal(t, int64(100_00), m.AmountNet(), "19: Net from net")
}

func TestMoneyAdd(t *testing.T) {
	testCases := []struct {
		m1          *money.Money
		m2          *money.Money
		amountNet   int64
		amountGross int64
		err         error
	}{
		{
			m1:          money.NewFromGross(100_00, "EUR", money.VAT_19_00),
			m2:          money.NewFromGross(200_00, "EUR", money.VAT_19_00),
			amountNet:   252_10,
			amountGross: 300_00,
		},
		{
			m1:          money.NewFromNet(100_00, "EUR", money.VAT_19_00),
			m2:          money.NewFromNet(200_00, "EUR", money.VAT_19_00),
			amountNet:   300_00,
			amountGross: 357_00,
		},
		{
			m1:          money.NewFromGross(333_33, "EUR", money.VAT_19_00),
			m2:          money.NewFromGross(712_59, "EUR", money.VAT_19_00),
			amountNet:   878_92,
			amountGross: 1045_92,
		},
		{
			m1:          money.NewFromNet(333_33, "EUR", money.VAT_19_00),
			m2:          money.NewFromNet(712_59, "EUR", money.VAT_19_00),
			amountNet:   1045_92,
			amountGross: 1244_64,
		},
		{
			m1:          money.NewFromGross(333_33, "EUR", money.VAT_07_00),
			m2:          money.NewFromGross(712_59, "EUR", money.VAT_07_00),
			amountNet:   977_50,
			amountGross: 1045_92,
		},
		{
			m1:          money.NewFromNet(333_33, "EUR", money.VAT_07_00),
			m2:          money.NewFromNet(712_59, "EUR", money.VAT_07_00),
			amountNet:   1045_92,
			amountGross: 1119_13,
		},
		{
			m1:  money.NewFromNet(333_33, "EUR", money.VAT_07_00),
			m2:  money.NewFromNet(712_59, "EUR", money.VAT_19_00),
			err: money.ErrCurrencyOrVatRateMissmatch,
		},
	}

	for _, testCase := range testCases {
		m3, err := testCase.m1.Add(testCase.m2)
		require.ErrorIs(t, err, testCase.err)
		if err == nil {
			require.Equal(t, testCase.amountNet, m3.AmountNet(), fmt.Sprintf("Test %s + %s", testCase.m1.Display(), testCase.m2.Display()))
			require.Equal(t, testCase.amountGross, m3.AmountGross(), fmt.Sprintf("Test %s + %s", testCase.m1.Display(), testCase.m2.Display()))
		}
	}
}

func TestMoneySubtract(t *testing.T) {
	testCases := []struct {
		m1          *money.Money
		m2          *money.Money
		amountNet   int64
		amountGross int64
		err         error
	}{
		{
			m1:          money.NewFromGross(100_00, "EUR", money.VAT_19_00),
			m2:          money.NewFromGross(200_00, "EUR", money.VAT_19_00),
			amountNet:   -84_03,
			amountGross: -100_00,
		},
		{
			m1:          money.NewFromNet(100_00, "EUR", money.VAT_19_00),
			m2:          money.NewFromNet(200_00, "EUR", money.VAT_19_00),
			amountNet:   -100_00,
			amountGross: -119_00,
		},
		{
			m1:          money.NewFromGross(333_33, "EUR", money.VAT_19_00),
			m2:          money.NewFromGross(712_59, "EUR", money.VAT_19_00),
			amountNet:   -318_71,
			amountGross: -379_26,
		},
		{
			m1:          money.NewFromNet(333_33, "EUR", money.VAT_19_00),
			m2:          money.NewFromNet(712_59, "EUR", money.VAT_19_00),
			amountNet:   -379_26,
			amountGross: -451_32,
		},
		{
			m1:          money.NewFromGross(333_33, "EUR", money.VAT_07_00),
			m2:          money.NewFromGross(712_59, "EUR", money.VAT_07_00),
			amountNet:   -354_45,
			amountGross: -379_26,
		},
		{
			m1:          money.NewFromNet(333_33, "EUR", money.VAT_07_00),
			m2:          money.NewFromNet(712_59, "EUR", money.VAT_07_00),
			amountNet:   -379_26,
			amountGross: -405_81,
		},
		{
			m1:  money.NewFromNet(333_33, "EUR", money.VAT_07_00),
			m2:  money.NewFromNet(712_59, "EUR", money.VAT_19_00),
			err: money.ErrCurrencyOrVatRateMissmatch,
		},
	}

	for _, testCase := range testCases {
		m3, err := testCase.m1.Subtract(testCase.m2)
		require.ErrorIs(t, err, testCase.err)
		if err == nil {
			require.Equal(t, testCase.amountNet, m3.AmountNet(), fmt.Sprintf("Test %s + %s", testCase.m1.Display(), testCase.m2.Display()))
			require.Equal(t, testCase.amountGross, m3.AmountGross(), fmt.Sprintf("Test %s + %s", testCase.m1.Display(), testCase.m2.Display()))
		}
	}
}

func TestMoneyMultiply(t *testing.T) {
	testCases := []struct {
		m1          *money.Money
		factors     []int64
		amountNet   []int64
		amountGross []int64
	}{
		{
			m1:          money.NewFromGross(100_00, "EUR", money.VAT_19_00),
			factors:     []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			amountNet:   []int64{0, 84_03, 168_07, 252_10, 336_13, 420_17, 504_20, 588_24, 672_27, 756_30, 840_34},
			amountGross: []int64{0, 100_00, 200_00, 300_00, 400_00, 500_00, 600_00, 700_00, 800_00, 900_00, 1000_00},
		},
		{
			m1:          money.NewFromNet(100_00, "EUR", money.VAT_19_00),
			factors:     []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			amountNet:   []int64{0, 100_00, 200_00, 300_00, 400_00, 500_00, 600_00, 700_00, 800_00, 900_00, 1000_00},
			amountGross: []int64{0, 119_00, 238_00, 357_00, 476_00, 595_00, 714_00, 833_00, 952_00, 1071_00, 1190_00},
		},
	}

	for _, testCase := range testCases {
		for idx, factor := range testCase.factors {
			m3 := testCase.m1.Multiply(factor)
			require.Equal(t, testCase.amountNet[idx], m3.AmountNet(), fmt.Sprintf("Test %s * %d", testCase.m1.Display(), factor))
			require.Equal(t, testCase.amountGross[idx], m3.AmountGross(), fmt.Sprintf("Test %s * %d", testCase.m1.Display(), factor))
		}
	}
}

func TestDisplay(t *testing.T) {
	testCases := []struct {
		m        *money.Money
		opts     []money.DisplayOption
		expected string
	}{
		{
			m:        money.NewFromGross(100_00, "EUR", money.VAT_19_00),
			expected: "€100.00",
		},
		{
			m:        money.NewFromGross(100_00, "EUR", money.VAT_19_00),
			expected: "€100.00 (Gross@19.00%)",
			opts: []money.DisplayOption{
				money.DisplayVatSuffix,
			},
		},
		{
			m:        money.NewFromNet(100_00, "EUR", money.VAT_19_00),
			expected: "€100.00",
		},
		{
			m:        money.NewFromNet(100_00, "EUR", money.VAT_19_00),
			expected: "€100.00 (Net@19.00%)",
			opts: []money.DisplayOption{
				money.DisplayVatSuffix,
			},
		},
	}
	for _, testCase := range testCases {
		require.Equal(t, testCase.expected, testCase.m.Display(testCase.opts...))
	}
}

func TestIsSomething(t *testing.T) {
	m := money.NewFromGross(0, "EUR", money.VAT_19_00)
	require.True(t, m.IsZero(), "0 is zero")
	require.False(t, m.IsPositive(), "0 is not positive")
	require.False(t, m.IsNegative(), "0 is not negative")
	m = money.NewFromGross(1_00, "EUR", money.VAT_19_00)
	require.False(t, m.IsZero(), "1_00 is not zero")
	require.True(t, m.IsPositive(), "1_00 is positive")
	require.False(t, m.IsNegative(), "1_00 is not negative")
	m = money.NewFromGross(-1_00, "EUR", money.VAT_19_00)
	require.False(t, m.IsZero(), "-1_00 is not zero")
	require.False(t, m.IsPositive(), "-1_00 is not positive")
	require.True(t, m.IsNegative(), "-1_00 is negative")
}

func TestCurrency(t *testing.T) {
	m := money.NewFromGross(0, "EUR", money.VAT_19_00)
	require.Equal(t, "EUR", m.Currency())
}

func TestCompare(t *testing.T) {
	testCases := []struct {
		m1                 *money.Money
		m2                 *money.Money
		equal              bool
		lessThan           bool
		lessThanOrEqual    bool
		greaterThan        bool
		greaterThanOrEqual bool
		compare            int
		err                error
	}{
		{
			m1:                 money.NewFromGross(100_00, "EUR", money.VAT_19_00),
			m2:                 money.NewFromGross(100_00, "EUR", money.VAT_19_00),
			equal:              true,
			lessThan:           false,
			lessThanOrEqual:    true,
			greaterThan:        false,
			greaterThanOrEqual: true,
			compare:            0,
		},
		{
			m1:                 money.NewFromGross(100_00, "EUR", money.VAT_19_00),
			m2:                 money.NewFromGross(150_00, "EUR", money.VAT_19_00),
			equal:              false,
			lessThan:           true,
			lessThanOrEqual:    true,
			greaterThan:        false,
			greaterThanOrEqual: false,
			compare:            -1,
		},
		{
			m1:                 money.NewFromGross(100_00, "EUR", money.VAT_19_00),
			m2:                 money.NewFromGross(50_00, "EUR", money.VAT_19_00),
			equal:              false,
			lessThan:           false,
			lessThanOrEqual:    false,
			greaterThan:        true,
			greaterThanOrEqual: true,
			compare:            1,
		},
		{
			m1:                 money.NewFromGross(100_00, "EUR", money.VAT_19_00),
			m2:                 money.NewFromGross(100_00, "USD", money.VAT_19_00),
			equal:              false,
			lessThan:           false,
			lessThanOrEqual:    false,
			greaterThan:        false,
			greaterThanOrEqual: false,
			compare:            0,
			err:                money.ErrCurrencyOrVatRateMissmatch,
		},
		{
			m1:                 money.NewFromGross(100_00, "EUR", money.VAT_19_00),
			m2:                 money.NewFromGross(100_00, "EUR", money.VAT_07_00),
			equal:              false,
			lessThan:           false,
			lessThanOrEqual:    false,
			greaterThan:        false,
			greaterThanOrEqual: false,
			compare:            0,
			err:                money.ErrCurrencyOrVatRateMissmatch,
		},
	}
	for _, testCase := range testCases {
		equal, err := testCase.m1.Equals(testCase.m2)
		require.ErrorIs(t, err, testCase.err, "error is")
		require.Equal(t, testCase.equal, equal, "is equal")
		lessThan, err := testCase.m1.LessThan(testCase.m2)
		require.ErrorIs(t, err, testCase.err, "error is")
		require.Equal(t, testCase.lessThan, lessThan, "less than")
		lessThanOrEqual, err := testCase.m1.LessThanOrEqual(testCase.m2)
		require.ErrorIs(t, err, testCase.err, "error is")
		require.Equal(t, testCase.lessThanOrEqual, lessThanOrEqual, "less than or equal")
		greaterThan, err := testCase.m1.GreaterThan(testCase.m2)
		require.ErrorIs(t, err, testCase.err, "error is")
		require.Equal(t, testCase.greaterThan, greaterThan, "greater than")
		greaterThanOrEqual, err := testCase.m1.GreaterThanOrEqual(testCase.m2)
		require.ErrorIs(t, err, testCase.err, "error is")
		require.Equal(t, testCase.greaterThanOrEqual, greaterThanOrEqual, "greater than or equal")
		compare, err := testCase.m1.Compare(testCase.m2)
		require.ErrorIs(t, err, testCase.err, "error is")
		require.Equal(t, testCase.compare, compare)
	}
}

func TestVatIncluded(t *testing.T) {
	testCases := []struct {
		m         *money.Money
		vatAmount int64
	}{
		{
			m:         money.NewFromGross(119_00, "EUR", money.VAT_19_00),
			vatAmount: 19_00,
		},
		{
			m:         money.NewFromNet(100_00, "EUR", money.VAT_19_00),
			vatAmount: 19_00,
		},
		{
			m:         money.NewFromGross(107_00, "EUR", money.VAT_07_00),
			vatAmount: 7_00,
		},
		{
			m:         money.NewFromNet(100_00, "EUR", money.VAT_07_00),
			vatAmount: 7_00,
		},
	}

	for _, testCase := range testCases {
		vat := testCase.m.VATIncluded()
		require.Equal(t, testCase.vatAmount, vat.Amount())
	}
}

func TestAbsoluteAndNegative(t *testing.T) {
	m1 := money.NewFromGross(100_00, "EUR", money.VAT_19_00)
	m2 := money.NewFromGross(-100_00, "EUR", money.VAT_19_00)

	require.True(t, must.Must(m1.Negative().Equals(m2)), "neg(m1) = m2")
	require.True(t, must.Must(m1.Equals(m2.Absolute())), "m1 = abs(m2)")
	require.True(t, must.Must(m1.Negative().Equals(m2.Negative())), "neg(m1) = neg(m2)")
}

func TestSplit(t *testing.T) {
	m := money.NewFromGross(100_00, "EUR", money.VAT_19_00)

	splitted, err := m.Split(4)
	require.NoError(t, err)
	require.Len(t, splitted, 4)
	for _, s := range splitted {
		require.True(t, must.Must(money.NewFromGross(25_00, "EUR", money.VAT_19_00).Equals(s)))
	}
}

func TestAllocate(t *testing.T) {
	m := money.NewFromGross(100_00, "EUR", money.VAT_19_00)

	allocated, err := m.Allocate(4)
	require.NoError(t, err)
	require.Len(t, allocated, 1)
	for _, s := range allocated {
		require.True(t, must.Must(money.NewFromGross(100_00, "EUR", money.VAT_19_00).Equals(s)))
	}

	allocated, err = m.Allocate(30, 40, 15, 15)
	require.NoError(t, err)
	require.Len(t, allocated, 4)

	require.True(t, must.Must(money.NewFromGross(30_00, "EUR", money.VAT_19_00).Equals(allocated[0])))
	require.True(t, must.Must(money.NewFromGross(40_00, "EUR", money.VAT_19_00).Equals(allocated[1])))
	require.True(t, must.Must(money.NewFromGross(15_00, "EUR", money.VAT_19_00).Equals(allocated[2])))
	require.True(t, must.Must(money.NewFromGross(15_00, "EUR", money.VAT_19_00).Equals(allocated[3])))
}

func TestAsMaijorUnits(t *testing.T) {
	m1 := money.NewFromGross(100_00, "EUR", money.VAT_19_00)

	require.Equal(t, 100.00, m1.AsMajorUnits())
}

func TestRound(t *testing.T) {
	m1 := money.NewFromGross(100_00, "EUR", money.VAT_19_00)
	m2 := money.NewFromGross(100_45, "EUR", money.VAT_19_00)
	require.True(t, must.Must(m1.Equals(m1.Round())))
	require.True(t, must.Must(m1.Equals(m2.Round())))
}

func TestNegate(t *testing.T) {
	m1 := money.NewFromGross(100_00, "EUR", money.VAT_19_00)
	m2 := money.NewFromGross(-100_00, "EUR", money.VAT_19_00)

	require.Equal(t, int64(-100_00), m1.Negate().AmountGross())
	require.Equal(t, int64(100_00), m2.Negate().AmountGross())
}
