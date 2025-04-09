package ct_test

import (
	"testing"

	"github.com/mansio-gmbh/goapiutils/ct"
	"github.com/mansio-gmbh/goapiutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestIsEqual(t *testing.T) {
	tests := []struct {
		unit1, unit2 ct.UnitValue
		expected     bool
	}{
		{ct.UnitValue{ptr.Val("m"), 10}, ct.UnitValue{ptr.Val("m"), 10}, true},
		{ct.UnitValue{ptr.Val("m"), 5}, ct.UnitValue{ptr.Val("m"), 10}, false},
		{ct.UnitValue{ptr.Val("m"), 5}, ct.UnitValue{ptr.Val("cm"), 500}, true},
		{ct.UnitValue{ptr.Val("m"), 5}, ct.UnitValue{ptr.Val("mi"), 0.00310686}, true},
		{ct.UnitValue{ptr.Val("m"), 5}, ct.UnitValue{ptr.Val("mile"), 0.00310686}, true},
		{ct.UnitValue{ptr.Val("cm"), 5}, ct.UnitValue{ptr.Val("mm"), 50}, true},
		{ct.UnitValue{ptr.Val("cm"), 5}, ct.UnitValue{ptr.Val("m"), 0.05}, true},
		{ct.UnitValue{ptr.Val("cm"), 5}, ct.UnitValue{ptr.Val("km"), 0.00005}, true},
		{ct.UnitValue{ptr.Val("cm"), 5}, ct.UnitValue{ptr.Val("ft"), 0.164042}, true},
		{ct.UnitValue{ptr.Val("cm"), 5}, ct.UnitValue{ptr.Val("yd"), 0.0546807}, true},
		{ct.UnitValue{ptr.Val("cm"), 5}, ct.UnitValue{ptr.Val("inch"), 1.9685}, true},
		{ct.UnitValue{ptr.Val("cm"), 5}, ct.UnitValue{ptr.Val("cm"), 10}, false},
		{ct.UnitValue{ptr.Val("g"), 10000}, ct.UnitValue{ptr.Val("kg"), 10}, true},
		{ct.UnitValue{ptr.Val("t"), 1}, ct.UnitValue{ptr.Val("kg"), 1000}, true},
		{ct.UnitValue{ptr.Val("kg"), 5}, ct.UnitValue{ptr.Val("kg"), 10}, false},
		{ct.UnitValue{ptr.Val("g"), 5}, ct.UnitValue{ptr.Val("g"), 10}, false},
	}

	for _, test := range tests {
		result := test.unit1.IsEqual(&test.unit2)
		require.Equal(t, test.expected, result, "Expected %v but got %v for unit1: %v and unit2: %v", test.expected, result, test.unit1, test.unit2)
	}
}

func TestAddUnitValues(t *testing.T) {
	tests := []struct {
		unit1    *ct.UnitValue
		units    []*ct.UnitValue
		expected *ct.UnitValue
	}{
		{&ct.UnitValue{ptr.Val("m"), 10}, []*ct.UnitValue{{ptr.Val("m"), 5}}, &ct.UnitValue{ptr.Val("m"), 15}},
		{&ct.UnitValue{ptr.Val("m"), 5}, []*ct.UnitValue{{ptr.Val("cm"), 500}}, &ct.UnitValue{ptr.Val("m"), 10}},
		{&ct.UnitValue{ptr.Val("cm"), 5}, []*ct.UnitValue{{ptr.Val("mm"), 50}}, &ct.UnitValue{ptr.Val("cm"), 10}},
		{&ct.UnitValue{ptr.Val("cm"), 5}, []*ct.UnitValue{{ptr.Val("mm"), 50}}, &ct.UnitValue{ptr.Val("cm"), 10}},
		{&ct.UnitValue{ptr.Val("m"), 5}, []*ct.UnitValue{{ptr.Val("cm"), 500}, {ptr.Val("mm"), 1000}}, &ct.UnitValue{ptr.Val("m"), 11}},
		{nil, []*ct.UnitValue{{ptr.Val("cm"), 500}, {ptr.Val("mm"), 1000}}, &ct.UnitValue{ptr.Val("cm"), 600}},
		{nil, []*ct.UnitValue{nil, nil, {ptr.Val("cm"), 500}, {ptr.Val("mm"), 1000}}, &ct.UnitValue{ptr.Val("cm"), 600}},
		{nil, []*ct.UnitValue{nil, nil}, nil},
	}

	for _, test := range tests {
		result, err := ct.AddUnitValues(test.unit1, test.units...)
		require.NoError(t, err)
		if test.expected == nil {
			require.Nil(t, result, "Expected nil but got %v for unit1: %v and units: %v", result, test.unit1, test.units)
		} else {
			require.Equal(t, *test.expected, *result, "Expected %v but got %v for unit1: %v and unit2: %v", test.expected, *result, test.unit1, test.units)
		}
	}
}
