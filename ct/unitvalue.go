package ct

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var (
	ErrUnitsIncompatible   = errors.New("units incompatible")
	ErrConversionUnitEmpty = errors.New("conversion target unit is empty")
)

type UnitValue struct {
	Unit  string  `json:"unit" dynamodbav:"unit"`
	Value float64 `json:"value" dynamodbav:"value"`
}

func NewUnitValue(value float64, unit string) *UnitValue {
	return &UnitValue{
		Unit:  unit,
		Value: value,
	}
}

const eps = 0.0001

func (u *UnitValue) IsEqual(other *UnitValue) bool {
	if other == nil {
		return false
	}
	if u.Unit == other.Unit && u.Value == other.Value {
		return true
	}

	convertedUV := UnitValue{
		Unit:  other.Unit,
		Value: other.Value,
	}
	err := convertedUV.ConvertUnit(u.Unit)
	if err != nil {
		return false
	}

	return math.Abs(convertedUV.Value-u.Value) < eps
}

func (u *UnitValue) IsValid() bool {
	return u.Unit != ""
}

func (u *UnitValue) IsZero() bool {
	return u.Value == 0
}

func (u *UnitValue) String() string {
	return fmt.Sprintf("%f%s", u.Value, u.Unit)
}

func (u *UnitValue) Add(other *UnitValue) error {
	if other == nil {
		return nil
	}
	if u.Unit == other.Unit {
		u.Value += other.Value
		return nil
	}

	tmpUV := UnitValue{
		Unit:  other.Unit,
		Value: other.Value,
	}
	err := tmpUV.ConvertUnit(u.Unit)
	if err != nil {
		return err
	}

	u.Value += tmpUV.Value

	return nil
}

func (u *UnitValue) ConvertUnit(targetUnit string) error {
	if targetUnit == "" {
		return ErrConversionUnitEmpty
	}

	knownConversions := map[string]map[string]float64{
		"m":    {"m": 1, "cm": 100, "mm": 1000, "km": 0.001, "ft": 3.28084, "mi": 0.000621371, "yd": 1.09361, "inch": 39.3701, "mile": 0.000621371},
		"cm":   {"m": 0.01, "cm": 1, "mm": 10, "km": 0.00001, "ft": 0.0328084, "mi": 0.0000062137, "yd": 0.0109361, "inch": 0.393701, "mile": 0.0000062137},
		"mm":   {"m": 0.001, "cm": 0.1, "mm": 1, "km": 0.000001, "ft": 0.00328084, "mi": 0.000000621371, "yd": 0.00328084, "inch": 0.0393701, "mile": 0.000000621371},
		"km":   {"m": 1000, "cm": 100000, "mm": 1000000, "km": 1, "ft": 3280.84, "mi": 0.621371, "yd": 1093.61, "inch": 39370.1, "mile": 0.621371},
		"ft":   {"m": 0.3048, "cm": 30.48, "mm": 304.8, "km": 0.0003048, "ft": 1, "mi": 0.000189394, "yd": 0.333333, "inch": 12, "mile": 0.000189394},
		"mi":   {"m": 1609.34, "cm": 160934, "mm": 1609340, "km": 1.60934, "ft": 5280, "mi": 1, "yd": 1760, "inch": 63360, "mile": 1},
		"yd":   {"m": 0.9144, "cm": 91.44, "mm": 914.4, "km": 0.0009144, "ft": 3, "mi": 0.000568182, "yd": 1, "inch": 36, "mile": 0.000568182},
		"inch": {"m": 0.0254, "cm": 2.54, "mm": 25.4, "km": 0.0000254, "ft": 0.0833333, "mi": 0.0000157828, "yd": 0.0277778, "inch": 1, "mile": 0.0000157828},
		"mile": {"m": 1609.34, "cm": 160934, "mm": 1609340, "km": 1.60934, "ft": 5280, "mi": 1, "yd": 1760, "inch": 63360, "mile": 1},
		"kg":   {"kg": 1, "g": 1000, "t": 0.001, "mg": 1000000, "lb": 2.20462, "oz": 35.274},
		"g":    {"kg": 0.001, "g": 1, "t": 0.000001, "mg": 1000, "lb": 0.00220462, "oz": 0.035274},
		"t":    {"kg": 1000, "g": 1000000, "t": 1, "mg": 1000000000, "lb": 2204.62, "oz": 35274},
		"mg":   {"kg": 0.000001, "g": 0.001, "t": 0.000000001, "mg": 1, "lb": 0.00000220462, "oz": 0.000035274},
		"lb":   {"kg": 0.453592, "g": 453.592, "t": 0.000453592, "mg": 453592, "lb": 1, "oz": 16},
		"oz":   {"kg": 0.0283495, "g": 28.3495, "t": 0.0000283495, "mg": 28349.5, "lb": 0.0625, "oz": 1},
		"l":    {"l": 1, "ml": 1000, "cl": 100, "dl": 10, "m3": 0.001, "gal": 0.264172, "qt": 1.05669, "pt": 2.11338, "m³": 0.001},
		"ml":   {"l": 0.001, "ml": 1, "cl": 0.1, "dl": 0.01, "m3": 0.000001, "gal": 0.000264172, "qt": 0.00105669, "pt": 0.00211338, "m³": 0.000001},
		"cl":   {"l": 0.01, "ml": 10, "cl": 1, "dl": 0.1, "m3": 0.00001, "gal": 0.00264172, "qt": 0.0105669, "pt": 0.0211338, "m³": 0.00001},
		"dl":   {"l": 0.1, "ml": 100, "cl": 10, "dl": 1, "m3": 0.0001, "gal": 0.0264172, "qt": 0.105669, "pt": 0.211338, "m³": 0.0001},
		"m3":   {"l": 1000, "ml": 1000000, "cl": 100000, "dl": 10000, "m3": 1, "gal": 264.172, "qt": 1056.69, "pt": 2113.38, "m³": 1},
		"gal":  {"l": 3.78541, "ml": 3785.41, "cl": 378.541, "dl": 37.8541, "m3": 0.00378541, "gal": 1, "qt": 4, "pt": 8, "m³": 0.00378541},
		"qt":   {"l": 0.946353, "ml": 946.353, "cl": 94.6353, "dl": 9.46353, "m3": 0.000946353, "gal": 0.25, "qt": 1, "pt": 2, "m³": 0.000946353},
		"pt":   {"l": 0.473176, "ml": 473.176, "cl": 47.3176, "dl": 4.73176, "m3": 0.000473176, "gal": 0.125, "qt": 0.5, "pt": 1, "m³": 0.000473176},
		"m³":   {"l": 1000, "ml": 1000000, "cl": 100000, "dl": 10000, "m3": 1, "gal": 264.172, "qt": 1056.69, "pt": 2113.38, "m³": 1},
		"m²":   {"m²": 1, "m2": 1, "cm²": 10000, "cm2": 10000, "mm²": 1000000, "mm3": 1000000, "km²": 0.000001, "km2": 0.000001, "ft²": 10.7639, "ft2": 10.7639, "in²": 1550.0031, "in2": 1550.0031, "ha": 0.0001, "ac": 0.000247105},
		"m2":   {"m²": 1, "m2": 1, "cm²": 10000, "cm2": 10000, "mm²": 1000000, "mm3": 1000000, "km²": 0.000001, "km2": 0.000001, "ft²": 10.7639, "ft2": 10.7639, "in²": 1550.0031, "in2": 1550.0031, "ha": 0.0001, "ac": 0.000247105},
		"cm²":  {"m²": 0.0001, "m2": 0.0001, "cm²": 1, "cm2": 1, "mm²": 100, "mm3": 10000, "km²": 0.000000001, "km2": 0.000000001, "ft²": 0.00107639, "ft2": 0.00107639, "in²": 0.15500031, "in2": 0.15500031, "ha": 0.00000001, "ac": 0.0000000247105},
		"cm2":  {"m²": 0.0001, "m2": 0.0001, "cm²": 1, "cm2": 1, "mm²": 100, "mm3": 10000, "km²": 0.000000001, "km2": 0.000000001, "ft²": 0.00107639, "ft2": 0.00107639, "in²": 0.15500031, "in2": 0.15500031, "ha": 0.00000001, "ac": 0.0000000247105},
		"mm²":  {"m²": 0.000001, "m2": 0.000001, "cm²": 0.01, "cm2": 0.01, "mm²": 1, "mm3": 10, "km²": 0.000000000001, "km2": 0.000000000001, "ft²": 0.0000107639, "ft2": 0.0000107639, "in²": 0.0015500031, "in2": 0.0015500031, "ha": 0.0000000001, "ac": 0.000000000247105},
		"km²":  {"m²": 1000000, "m2": 1000000, "cm²": 10000000000, "cm2": 10000000000, "mm²": 1000000000000, "mm3": 1000000000000, "km²": 1, "km2": 1, "ft²": 10763910.4, "ft2": 10763910.4, "in²": 1550003100.01, "in2": 1550003100.01, "ha": 100, "ac": 247.105},
		"ft²":  {"m²": 0.092903, "m2": 0.092903, "cm²": 929.0304, "cm2": 929.0304, "mm²": 92903.04, "mm3": 92903.04, "km²": 0.000000092903, "km2": 0.000000092903, "ft²": 1, "ft2": 1, "in²": 144, "in2": 144, "ha": 0.0000092903, "ac": 0.0000229568},
		"ft2":  {"m²": 0.092903, "m2": 0.092903, "cm²": 929.0304, "cm2": 929.0304, "mm²": 92903.04, "mm3": 92903.04, "km²": 0.000000092903, "km2": 0.000000092903, "ft²": 1, "ft2": 1, "in²": 144, "in2": 144, "ha": 0.0000092903, "ac": 0.0000229568},
		"in²":  {"m²": 0.00064516, "m2": 0.00064516, "cm²": 6.4516, "cm2": 6.4516, "mm²": 645.16, "mm3": 645.16, "km²": 0.00000000064516, "km2": 0.00000000064516, "ft²": 0.00694444, "ft2": 0.00694444, "in²": 1, "in2": 1, "ha": 0.000000064516, "ac": 0.000000159422},
		"in2":  {"m²": 0.00064516, "m2": 0.00064516, "cm²": 6.4516, "cm2": 6.4516, "mm²": 645.16, "mm3": 645.16, "km²": 0.00000000064516, "km2": 0.00000000064516, "ft²": 0.00694444, "ft2": 0.00694444, "in²": 1, "in2": 1, "ha": 0.000000064516, "ac": 0.000000159422},
		"ha":   {"m²": 10000, "m2": 10000, "cm²": 100000000, "cm2": 100000000, "mm²": 1000000000000, "mm3": 1000000000000, "km²": 0.01, "km2": 0.01, "ft²": 107639.104, "ft2": 107639.104, "in²": 15500031.0, "in2": 15500031.0, "ha": 1, "ac": 2.47105},
		"ac":   {"m²": 4046.86, "m2": 4046.86, "cm²": 40468600, "cm2": 40468600, "mm²": 4046860000000, "mm3": 4046860000000, "km²": 0.00404686, "km2": 0.00404686, "ft²": 43560, "ft2": 43560, "in²": 62726400, "in2": 62726400, "ha": 0.404686, "ac": 1},
		"lm":   {"lm": 1},
		"kw":   {"kw": 1, "w": 1000, "hp": 1.34102, "ps": 1.35962},
		"w":    {"kw": 0.001, "w": 1, "hp": 0.00134102, "ps": 0.00135962},
		"hp":   {"kw": 0.7457, "w": 745.7, "hp": 1, "ps": 1.01442},
		"ps":   {"kw": 0.7355, "w": 735.5, "hp": 0.98632, "ps": 1},
	}

	sourceUnit := strings.ToLower(u.Unit)
	targetUnit = strings.ToLower(targetUnit)

	if sourceUnit == targetUnit {
		return nil
	}

	conversion, okConversion := knownConversions[sourceUnit]
	if !okConversion {
		return ErrUnitsIncompatible
	}
	factor, okFactor := conversion[targetUnit]
	if !okFactor {
		return ErrUnitsIncompatible
	}

	u.Value = u.Value * factor

	return nil
}
