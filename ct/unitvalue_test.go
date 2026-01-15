package ct_test

import (
	"testing"

	"github.com/mansio-gmbh/goapiutils/ct"
	"github.com/stretchr/testify/require"
)

func TestIsEqual(t *testing.T) {
	tests := []struct {
		unit1, unit2 ct.UnitValue
		expected     bool
	}{
		{ct.UnitValue{"m", 10}, ct.UnitValue{"m", 10}, true},
		{ct.UnitValue{"m", 5}, ct.UnitValue{"m", 10}, false},
		{ct.UnitValue{"m", 5}, ct.UnitValue{"cm", 500}, true},
		{ct.UnitValue{"m", 5}, ct.UnitValue{"mi", 0.00310686}, true},
		{ct.UnitValue{"m", 5}, ct.UnitValue{"mile", 0.00310686}, true},
		{ct.UnitValue{"cm", 5}, ct.UnitValue{"mm", 50}, true},
		{ct.UnitValue{"cm", 5}, ct.UnitValue{"m", 0.05}, true},
		{ct.UnitValue{"cm", 5}, ct.UnitValue{"km", 0.00005}, true},
		{ct.UnitValue{"cm", 5}, ct.UnitValue{"ft", 0.164042}, true},
		{ct.UnitValue{"cm", 5}, ct.UnitValue{"yd", 0.0546807}, true},
		{ct.UnitValue{"cm", 5}, ct.UnitValue{"inch", 1.9685}, true},
		{ct.UnitValue{"cm", 5}, ct.UnitValue{"cm", 10}, false},
		{ct.UnitValue{"g", 10000}, ct.UnitValue{"kg", 10}, true},
		{ct.UnitValue{"t", 1}, ct.UnitValue{"kg", 1000}, true},
		{ct.UnitValue{"kg", 5}, ct.UnitValue{"kg", 10}, false},
		{ct.UnitValue{"g", 5}, ct.UnitValue{"g", 10}, false},
	}

	for _, test := range tests {
		result := test.unit1.IsEqual(&test.unit2)
		require.Equal(t, test.expected, result, "Expected %v but got %v for unit1: %v and unit2: %v", test.expected, result, test.unit1, test.unit2)
	}
}
