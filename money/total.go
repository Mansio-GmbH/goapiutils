package money

import (
	"github.com/elliotchance/pie/v2"
	"github.com/mansio-gmbh/goapiutils/must"
)

type Total struct {
	netTotal   map[string]*MoneyWithoutVat
	grossTotal map[string]*MoneyWithoutVat
	vatTotal   map[string]*MoneyWithoutVat
	vatByCode  map[string]map[string]*MoneyWithoutVat
}

func NewTotal(initialValues ...*Money) *Total {
	t := &Total{
		netTotal:   make(map[string]*MoneyWithoutVat),
		grossTotal: make(map[string]*MoneyWithoutVat),
		vatTotal:   make(map[string]*MoneyWithoutVat),
		vatByCode:  make(map[string]map[string]*MoneyWithoutVat),
	}
	for _, m := range initialValues {
		t.Add(m)
	}
	return t
}

func NewTotal2(initialValues ...*Total) *Total {
	t := &Total{
		netTotal:   make(map[string]*MoneyWithoutVat),
		grossTotal: make(map[string]*MoneyWithoutVat),
		vatTotal:   make(map[string]*MoneyWithoutVat),
		vatByCode:  make(map[string]map[string]*MoneyWithoutVat),
	}
	for _, m := range initialValues {
		t = t.AddTotal(m)
	}
	return t
}

func (t *Total) Add(m *Money) {
	currencyCode := m.Currency()
	vatCode := m.VAT()
	if _, found := t.netTotal[currencyCode]; !found {
		t.netTotal[currencyCode] = NewWithoutVat(0, currencyCode)
	}
	if _, found := t.grossTotal[currencyCode]; !found {
		t.grossTotal[currencyCode] = NewWithoutVat(0, currencyCode)
	}
	if _, found := t.vatTotal[currencyCode]; !found {
		t.vatTotal[currencyCode] = NewWithoutVat(0, currencyCode)
	}
	if _, found := t.vatByCode[vatCode]; !found {
		t.vatByCode[vatCode] = make(map[string]*MoneyWithoutVat)
	}
	if _, found := t.vatByCode[vatCode][currencyCode]; !found {
		t.vatByCode[vatCode][currencyCode] = NewWithoutVat(0, currencyCode)
	}

	t.netTotal[currencyCode] = must.Must(t.netTotal[currencyCode].Add(NewWithoutVatFromMoney(m.netMoney)))
	t.grossTotal[currencyCode] = must.Must(t.grossTotal[currencyCode].Add(NewWithoutVatFromMoney(m.grossMoney)))
	t.vatTotal[currencyCode] = must.Must(t.vatTotal[currencyCode].Add(m.VATIncluded()))
	t.vatByCode[vatCode][currencyCode] = must.Must(t.vatByCode[vatCode][currencyCode].Add(m.VATIncluded()))
}

func (t Total) NetTotal(currencyCode string) (*MoneyWithoutVat, bool) {
	money, found := t.netTotal[currencyCode]
	return money, found
}

func (t Total) NetTotalOrZero(currencyCode string) *MoneyWithoutVat {
	if money, found := t.netTotal[currencyCode]; found {
		return money
	}
	return NewWithoutVat(0, currencyCode)
}

func (t Total) NetTotals() []*MoneyWithoutVat {
	if len(t.netTotal) == 0 {
		return []*MoneyWithoutVat{}
	}
	return pie.Values(t.netTotal)
}

func (t Total) GrossTotal(currencyCode string) (*MoneyWithoutVat, bool) {
	money, found := t.grossTotal[currencyCode]
	return money, found
}

func (t Total) GrossTotalOrZero(currencyCode string) *MoneyWithoutVat {
	if money, found := t.grossTotal[currencyCode]; found {
		return money
	}
	return NewWithoutVat(0, currencyCode)
}

func (t Total) GrossTotals() []*MoneyWithoutVat {
	if len(t.grossTotal) == 0 {
		return []*MoneyWithoutVat{}
	}
	return pie.Values(t.grossTotal)
}

func (t Total) VatTotal(currencyCode string) (*MoneyWithoutVat, bool) {
	money, found := t.vatTotal[currencyCode]
	return money, found
}

func (t Total) VatTotalOrZero(currencyCode string) *MoneyWithoutVat {
	if money, found := t.vatTotal[currencyCode]; found {
		return money
	}
	return NewWithoutVat(0, currencyCode)
}

func (t Total) VatTotals() []*MoneyWithoutVat {
	if len(t.vatTotal) == 0 {
		return []*MoneyWithoutVat{}
	}
	return pie.Values(t.vatTotal)
}

func (t Total) VatTotalsByCode(vatCode string) ([]*MoneyWithoutVat, bool) {
	moneys, found := t.vatByCode[vatCode]
	if !found {
		return nil, false
	}
	return pie.Values(moneys), true
}

func (t Total) VatTotalByCode(currencyCode, vatCode string) (*MoneyWithoutVat, bool) {
	moneys, found := t.vatByCode[vatCode]
	if !found {
		return nil, false
	}
	money, found := moneys[currencyCode]
	if !found {
		return nil, false
	}
	return money, true
}

func (t Total) VatTotalByCodeOrZero(currencyCode, vatCode string) *MoneyWithoutVat {
	if money, found := t.VatTotalByCode(currencyCode, vatCode); found {
		return money
	}
	return NewWithoutVat(0, currencyCode)
}

func (t Total) AddTotal(ot *Total) *Total {
	nt := NewTotal()
	for currency, money := range t.netTotal {
		nt.netTotal[currency], _ = money.Add(ot.netTotal[currency])
	}
	for currency, money := range ot.netTotal {
		nt.netTotal[currency], _ = money.Add(t.netTotal[currency])
	}
	for currency, money := range t.grossTotal {
		nt.grossTotal[currency], _ = money.Add(ot.grossTotal[currency])
	}
	for currency, money := range ot.grossTotal {
		nt.grossTotal[currency], _ = money.Add(t.grossTotal[currency])
	}
	for currency, money := range t.vatTotal {
		nt.vatTotal[currency], _ = money.Add(ot.vatTotal[currency])
	}
	for currency, money := range ot.vatTotal {
		nt.vatTotal[currency], _ = money.Add(t.vatTotal[currency])
	}
	for vatCode, moneys := range t.vatByCode {
		if _, found := nt.vatByCode[vatCode]; !found {
			nt.vatByCode[vatCode] = map[string]*MoneyWithoutVat{}
		}
		for currency, money := range moneys {
			if _, found := ot.vatByCode[vatCode]; !found {
				nt.vatByCode[vatCode][currency], _ = money.Add(NewWithoutVat(0, currency))
			} else if omoney, found := ot.vatByCode[vatCode][currency]; !found {
				nt.vatByCode[vatCode][currency], _ = money.Add(NewWithoutVat(0, currency))
			} else {
				nt.vatByCode[vatCode][currency], _ = money.Add(omoney)
			}
		}
	}
	for vatCode, moneys := range ot.vatByCode {
		if _, found := nt.vatByCode[vatCode]; !found {
			nt.vatByCode[vatCode] = map[string]*MoneyWithoutVat{}
		}
		for currency, money := range moneys {
			if _, found := t.vatByCode[vatCode]; !found {
				nt.vatByCode[vatCode][currency], _ = money.Add(NewWithoutVat(0, currency))
			} else if omoney, found := t.vatByCode[vatCode][currency]; !found {
				nt.vatByCode[vatCode][currency], _ = money.Add(NewWithoutVat(0, currency))
			} else {
				nt.vatByCode[vatCode][currency], _ = money.Add(omoney)
			}
		}
	}

	return nt
}

func (t Total) Negate() *Total {
	nt := NewTotal()
	for currency, money := range t.netTotal {
		nt.netTotal[currency] = NewWithoutVat(-money.Amount(), currency)
	}
	for currency, money := range t.grossTotal {
		nt.grossTotal[currency] = NewWithoutVat(-money.Amount(), currency)
	}
	for currency, money := range t.vatTotal {
		nt.vatTotal[currency] = NewWithoutVat(-money.Amount(), currency)
	}
	for vatCode, moneys := range t.vatByCode {
		if _, found := nt.vatByCode[vatCode]; !found {
			nt.vatByCode[vatCode] = map[string]*MoneyWithoutVat{}
		}
		for currency, money := range moneys {
			nt.vatByCode[vatCode][currency] = NewWithoutVat(-money.Amount(), currency)
		}
	}
	return nt
}
