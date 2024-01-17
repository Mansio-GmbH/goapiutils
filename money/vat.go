package money

import "errors"

type vat struct {
	rate        int64
	denominator int64
	code        string
}

var (
	VAT_00_00 = &vat{rate: 0, denominator: 1, code: "VAT_00_00"}
	VAT_19_00 = &vat{rate: 19_00, denominator: 100_00, code: "VAT_19_00"}
	VAT_07_00 = &vat{rate: 7_00, denominator: 100_00, code: "VAT_07_00"}

	vatRates = map[string]*vat{
		"VAT_00_00": VAT_00_00,
		"VAT_19_00": VAT_19_00,
		"VAT_07_00": VAT_07_00,
	}
)

func (v *vat) SameRate(ovat *vat) bool {
	return v.rate == ovat.rate
}

func VatByCode(code string) (*vat, error) {
	vat, found := vatRates[code]
	if !found {
		return nil, errors.New("vat rate not found")
	}
	return vat, nil
}

func Display(code string) string {
	switch code {
	case VAT_00_00.code:
		return ""
	case VAT_19_00.code:
		return "19 % MwSt"
	case VAT_07_00.code:
		return "7 % MwSt"
	default:
		return "Unknown VAT"
	}
}
