package money

import (
	"github.com/elliotchance/pie/v2"
	"github.com/mansio-gmbh/goapiutils/must"
)

type vatcode = string
type currencycode = string

type Total struct {
	net map[currencycode]map[vatcode]*MoneyWithoutVat
}

func NewTotal(initialValues ...*Money) *Total {
	t := &Total{
		net: make(map[currencycode]map[vatcode]*MoneyWithoutVat),
	}
	for _, m := range initialValues {
		t.Add(m)
	}
	return t
}

func NewTotal2(initialValues ...*Total) *Total {
	t := &Total{
		net: make(map[currencycode]map[vatcode]*MoneyWithoutVat),
	}
	for _, m := range initialValues {
		t = t.AddTotal(m)
	}
	return t
}

func (t *Total) Add(m *Money) {
	currencyCode := m.Currency()
	vatCode := m.VAT()
	if _, found := t.net[currencyCode]; !found {
		t.net[currencyCode] = make(map[vatcode]*MoneyWithoutVat)
	}
	if _, found := t.net[currencyCode][vatCode]; !found {
		t.net[currencyCode][vatCode] = NewWithoutVat(0, currencyCode)
	}

	t.net[currencyCode][vatCode] = must.Must(t.net[currencyCode][vatCode].Add(NewWithoutVatFromMoney(m.netMoney)))
}

func (t Total) NetTotal(currencyCode string) (*MoneyWithoutVat, bool) {
	if _, found := t.net[currencyCode]; !found {
		return nil, found
	}
	money := NewWithoutVat(0, currencyCode)
	for _, m := range t.net[currencyCode] {
		money, _ = money.Add(m)
	}
	return money, true
}

func (t Total) NetTotalOrZero(currencyCode string) *MoneyWithoutVat {
	if money, found := t.NetTotal(currencyCode); found {
		return money
	}
	return NewWithoutVat(0, currencyCode)
}

func (t Total) NetTotals() []*MoneyWithoutVat {
	totals := make([]*MoneyWithoutVat, 0, len(t.net))
	for currency := range t.net {
		totals = append(totals, t.NetTotalOrZero(currency))
	}
	return totals
}

func (t Total) GrossTotal(currencyCode string) (*MoneyWithoutVat, bool) {
	if _, found := t.net[currencyCode]; !found {
		return nil, found
	}
	money := NewWithoutVat(0, currencyCode)
	for vatCode, m := range t.net[currencyCode] {
		vat, _ := VatByCode(vatCode)
		money, _ = money.Add(m.MultiplyByFloat(1.0 + vat.Multiplier()))
	}
	return money, true
}

func (t Total) GrossTotalOrZero(currencyCode string) *MoneyWithoutVat {
	if money, found := t.GrossTotal(currencyCode); found {
		return money
	}
	return NewWithoutVat(0, currencyCode)
}

func (t Total) GrossTotals() []*MoneyWithoutVat {
	totals := make([]*MoneyWithoutVat, 0, len(t.net))
	for currency := range t.net {
		totals = append(totals, t.GrossTotalOrZero(currency))
	}
	return totals
}

func (t Total) VatTotal(currencyCode string) (*MoneyWithoutVat, bool) {
	if _, found := t.net[currencyCode]; !found {
		return nil, found
	}
	total := NewWithoutVat(0, currencyCode)
	for vatCode, m := range t.net[currencyCode] {
		vat := must.Must(VatByCode(vatCode))
		total, _ = total.Add(m.MultiplyByFloat(vat.Multiplier()))
	}
	return total, true
}

func (t Total) VatTotalOrZero(currencyCode string) *MoneyWithoutVat {
	if money, found := t.VatTotal(currencyCode); found {
		return money
	}
	return NewWithoutVat(0, currencyCode)
}

func (t Total) VatTotals() []*MoneyWithoutVat {
	totals := make([]*MoneyWithoutVat, 0, len(t.net))
	for currency := range t.net {
		totals = append(totals, t.VatTotalOrZero(currency))
	}
	return totals
}

func (t Total) VatTotalsByCode(vatCode string) ([]*MoneyWithoutVat, bool) {
	vats := make([]*MoneyWithoutVat, 0, len(t.net))
	for _, moneys := range t.net {
		if money, found := moneys[vatCode]; found {
			vats = append(vats, money)
		}
	}
	return vats, len(vats) > 0
}

func (t Total) VatTotalByCode(currencyCode, vatCode string) (*MoneyWithoutVat, bool) {
	moneys, found := t.net[currencyCode]
	if !found {
		return nil, false
	}
	money, found := moneys[vatCode]
	if !found {
		return nil, false
	}
	vat := must.Must(VatByCode(vatCode))
	return money.MultiplyByFloat(vat.Multiplier()), true
}

func (t Total) VatTotalByCodeOrZero(currencyCode, vatCode string) *MoneyWithoutVat {
	if money, found := t.VatTotalByCode(currencyCode, vatCode); found {
		return money
	}
	return NewWithoutVat(0, currencyCode)
}

func (t Total) AddTotal(ot *Total) *Total {
	nt := NewTotal()
	for currency, moneyByVat := range t.net {
		if _, found := nt.net[currency]; !found {
			nt.net[currency] = make(map[vatcode]*MoneyWithoutVat)
		}
		for vatCode, money := range moneyByVat {
			if _, found := ot.net[currency][vatCode]; !found {
				nt.net[currency][vatCode], _ = money.Add(NewWithoutVat(0, currency))
			}
			nt.net[currency][vatCode], _ = money.Add(ot.net[currency][vatCode])
		}
	}
	for currency, moneyByVat := range ot.net {
		if _, found := nt.net[currency]; !found {
			nt.net[currency] = make(map[vatcode]*MoneyWithoutVat)
		}
		for vatCode, money := range moneyByVat {
			if _, found := t.net[currency][vatCode]; !found {
				nt.net[currency][vatCode], _ = money.Add(NewWithoutVat(0, currency))
			}
			nt.net[currency][vatCode], _ = money.Add(t.net[currency][vatCode])
		}
	}
	return nt
}

func (t Total) Negate() *Total {
	nt := NewTotal()
	for currency, moneyByVat := range t.net {
		if _, found := nt.net[currency]; !found {
			nt.net[currency] = make(map[vatcode]*MoneyWithoutVat)
		}
		for vatCode, money := range moneyByVat {
			nt.net[currency][vatCode] = NewWithoutVat(-money.Amount(), currency)
		}
	}
	return nt
}

func (t Total) Negative() *Total {
	nt := NewTotal()
	for currency, moneyByVat := range t.net {
		if _, found := nt.net[currency]; !found {
			nt.net[currency] = make(map[vatcode]*MoneyWithoutVat)
		}
		for vatCode, money := range moneyByVat {
			nt.net[currency][vatCode] = NewWithoutVat(money.Negative().Amount(), currency)
		}
	}
	return nt
}

func (t Total) Absolute() *Total {
	nt := NewTotal()
	for currency, moneyByVat := range t.net {
		if _, found := nt.net[currency]; !found {
			nt.net[currency] = make(map[vatcode]*MoneyWithoutVat)
		}
		for vatCode, money := range moneyByVat {
			nt.net[currency][vatCode] = NewWithoutVat(money.Absolute().Amount(), currency)
		}
	}
	return nt
}

func (t Total) IsZero() bool {
	for _, moneyByVat := range t.net {
		for _, money := range moneyByVat {
			if !money.IsZero() {
				return false
			}
		}
	}
	return true
}

func (t Total) IsNegative() bool {
	if len(t.net) == 0 {
		return false
	}
	for _, moneyByVat := range t.net {
		for _, money := range moneyByVat {
			if !money.IsNegative() {
				return false
			}
		}
	}
	return true
}

func (t Total) IsPositive() bool {
	if len(t.net) == 0 {
		return false
	}
	for _, moneyByVat := range t.net {
		for _, money := range moneyByVat {
			if !money.IsPositive() {
				return false
			}
		}
	}
	return true
}

func (t Total) CurrencyCodes() []string {
	return pie.Keys(t.net)
}

func (t Total) VatCodes() []string {
	vatCodes := make([]string, 0)
	for _, moneys := range t.net {
		vatCodes = append(vatCodes, pie.Keys(moneys)...)
	}
	return pie.Unique(vatCodes)
}

func (t Total) Multiply(n int64) *Total {
	nt := NewTotal()
	for currency, moneyByVat := range t.net {
		if _, found := nt.net[currency]; !found {
			nt.net[currency] = make(map[vatcode]*MoneyWithoutVat)
		}
		for vatCode, money := range moneyByVat {
			nt.net[currency][vatCode] = money.Multiply(n)
		}
	}
	return nt
}

func (t Total) MultiplyByFloat(x float64) *Total {
	nt := NewTotal()
	for currency, moneyByVat := range t.net {
		if _, found := nt.net[currency]; !found {
			nt.net[currency] = make(map[vatcode]*MoneyWithoutVat)
		}
		for vatCode, money := range moneyByVat {
			nt.net[currency][vatCode] = money.MultiplyByFloat(x)
		}
	}
	return nt
}

func (t Total) Percentage(perc float64) *Total {
	return t.MultiplyByFloat(perc / 100.0)
}

func (t Total) Equals(ot *Total) bool {

	isZero := func(moneys map[vatcode]*MoneyWithoutVat) bool {
		for _, money := range moneys {
			if !money.IsZero() {
				return false
			}
		}
		return true
	}

	for currency, moneyByVat := range t.net {
		if otMoneyByVat, found := ot.net[currency]; !found && !isZero(moneyByVat) {
			return false
		} else {
			for vatCode, money := range moneyByVat {
				if omoney, found := otMoneyByVat[vatCode]; !found || !must.Must(money.Equals(omoney)) {
					return false
				}
			}
		}
	}

	for currency, moneyByVat := range ot.net {
		if tMoneyByVat, found := t.net[currency]; !found && !isZero(moneyByVat) {
			return false
		} else {
			for vatCode, money := range moneyByVat {
				if omoney, found := tMoneyByVat[vatCode]; !found || !must.Must(money.Equals(omoney)) {
					return false
				}
			}
		}
	}

	return true
}
