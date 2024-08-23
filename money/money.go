package money

import (
	"errors"
	"fmt"
	"math"

	"github.com/Rhymond/go-money"
	"github.com/mansio-gmbh/goapiutils/must"
)

type Money struct {
	grossMoney          *money.Money
	netMoney            *money.Money
	vat                 *vat
	LeadingValueIsGross bool
}

type Currency = money.Currency

type displayOptions struct {
	DisplayVatSuffix     bool
	NetVatSuffixFormat   string
	GrossVatSuffixFormat string
}
type DisplayOption func(o *displayOptions)

func DisplayVatSuffix(o *displayOptions) {
	o.DisplayVatSuffix = true
	o.NetVatSuffixFormat = "%s (Net@%.2f%%)"
	o.GrossVatSuffixFormat = "%s (Gross@%.2f%%)"
}

var ErrCurrencyOrVatRateMissmatch = errors.New("currencies or vat rate does not match")

func New(amount int64, currency string, vat *vat, valueIsGross bool) *Money {
	return newWithMoney(money.New(amount, currency), vat, valueIsGross)
}

func NewFromFloat(amount float64, currency string, vat *vat, valueIsGross bool) *Money {
	currencyDecimals := math.Pow10(money.GetCurrency(currency).Fraction)
	amountCents := int64(math.Round(amount * currencyDecimals))
	return New(amountCents, currency, vat, valueIsGross)
}

func NewFromNet(amount int64, currency string, vat *vat) *Money {
	return newWithMoney(money.New(amount, currency), vat, false)
}

func NewFromGross(amount int64, currency string, vat *vat) *Money {
	return newWithMoney(money.New(amount, currency), vat, true)
}

func newWithMoney(m *money.Money, vat *vat, valueIsGross bool) *Money {
	var mon *Money
	if valueIsGross {
		mon = &Money{
			grossMoney:          m,
			vat:                 vat,
			LeadingValueIsGross: true,
		}
	} else {
		mon = &Money{
			netMoney:            m,
			vat:                 vat,
			LeadingValueIsGross: false,
		}
	}
	return mon.updateOtherValue()
}

func (m *Money) assertSameCurrencyAndVatRate(om *Money) error {
	if m.netMoney.SameCurrency(om.netMoney) && m.vat.SameRate(om.vat) {
		return nil
	}
	return ErrCurrencyOrVatRateMissmatch
}

func (m Money) AmountGross() int64 {
	return m.grossMoney.Amount()
}

func (m Money) Amount() int64 {
	return m.calculateWithMoney().Amount()
}

func (m Money) AmountNet() int64 {
	return m.netMoney.Amount()
}

func (m Money) AmountGrossAsMajorUnits() float64 {
	return m.grossMoney.AsMajorUnits()
}

func (m Money) AmountNetAsMajorUnits() float64 {
	return m.netMoney.AsMajorUnits()
}

func (m *Money) Currency() string {
	return m.calculateWithMoney().Currency().Code
}

func (m *Money) VAT() string {
	return m.vat.code
}

func (m *Money) Equals(om *Money) (bool, error) {
	if err := m.assertSameCurrencyAndVatRate(om); err != nil {
		return false, err
	}
	return must.Must(m.calculateWithMoney().Equals(om.netOrGross(m.LeadingValueIsGross))), nil
}

func (m *Money) GreaterThan(om *Money) (bool, error) {
	if err := m.assertSameCurrencyAndVatRate(om); err != nil {
		return false, err
	}
	return must.Must(m.calculateWithMoney().GreaterThan(om.netOrGross(m.LeadingValueIsGross))), nil
}

func (m *Money) GreaterThanOrEqual(om *Money) (bool, error) {
	if err := m.assertSameCurrencyAndVatRate(om); err != nil {
		return false, err
	}
	return must.Must(m.calculateWithMoney().GreaterThanOrEqual(om.netOrGross(m.LeadingValueIsGross))), nil
}

func (m *Money) LessThan(om *Money) (bool, error) {
	if err := m.assertSameCurrencyAndVatRate(om); err != nil {
		return false, err
	}
	return must.Must(m.calculateWithMoney().LessThan(om.netOrGross(m.LeadingValueIsGross))), nil
}

func (m *Money) LessThanOrEqual(om *Money) (bool, error) {
	if err := m.assertSameCurrencyAndVatRate(om); err != nil {
		return false, err
	}
	return must.Must(m.calculateWithMoney().LessThanOrEqual(om.netOrGross(m.LeadingValueIsGross))), nil
}

func (m *Money) IsZero() bool {
	return m.calculateWithMoney().IsZero()
}

func (m *Money) IsPositive() bool {
	return m.calculateWithMoney().IsPositive()
}

func (m *Money) IsNegative() bool {
	return m.calculateWithMoney().IsNegative()
}

func (m *Money) Absolute() *Money {
	return m.newWithAmount(m.calculateWithMoney().Absolute())
}

func (m *Money) Negative() *Money {
	return m.newWithAmount(m.calculateWithMoney().Negative())
}

func (m *Money) Negate() *Money {
	if m.IsNegative() {
		return m.Absolute()
	} else {
		return m.Negative()
	}
}

func (m *Money) Add(other *Money) (*Money, error) {
	if err := m.assertSameCurrencyAndVatRate(other); err != nil {
		return nil, err
	}

	return m.newWithAmount(must.Must(m.calculateWithMoney().Add(other.netOrGross(m.LeadingValueIsGross)))), nil
}

func (m *Money) Subtract(other *Money) (*Money, error) {
	if err := m.assertSameCurrencyAndVatRate(other); err != nil {
		return nil, err
	}

	return m.newWithAmount(must.Must(m.calculateWithMoney().Subtract(other.netOrGross(m.LeadingValueIsGross)))), nil
}

func (m *Money) Round() *Money {
	return m.newWithAmount(m.calculateWithMoney().Round())
}

func (m *Money) Multiply(n int64) *Money {
	return m.newWithAmount(m.calculateWithMoney().Multiply(n))
}

func (m *Money) MultiplyByFloat(x float64) *Money {
	calculateMoney := m.calculateWithMoney()
	newAmount := int64(math.Round(float64(calculateMoney.Amount()) * x))
	return m.newWithAmount(money.New(newAmount, calculateMoney.Currency().Code))
}

func (m *Money) Percentage(perc float64) *Money {
	return m.MultiplyByFloat(perc / 100.0)
}

func (m *Money) VATIncluded() *MoneyWithoutVat {
	money, _ := m.grossMoney.Subtract(m.netMoney)
	return &MoneyWithoutVat{money: *money}
}

func (m *Money) Split(n int) ([]*Money, error) {
	splits, err := m.calculateWithMoney().Split(n)
	if err != nil {
		return nil, err
	}
	result := make([]*Money, len(splits))
	for idx, split := range splits {
		result[idx] = m.newWithAmount(split)
	}
	return result, nil
}

func (m *Money) Allocate(rs ...int) ([]*Money, error) {
	allocs, err := m.calculateWithMoney().Allocate(rs...)
	if err != nil {
		return nil, err
	}
	result := make([]*Money, len(allocs))
	for idx, alloc := range allocs {
		result[idx] = m.newWithAmount(alloc)
	}
	return result, nil
}

func (m *Money) AsMajorUnits() float64 {
	return m.calculateWithMoney().AsMajorUnits()
}

func (m *Money) Compare(om *Money) (int, error) {
	if err := m.assertSameCurrencyAndVatRate(om); err != nil {
		return 0, err
	}
	return m.calculateWithMoney().Compare(om.netOrGross(m.LeadingValueIsGross))
}

func (m *Money) updateOtherValue() *Money {
	if m.LeadingValueIsGross {
		amount := float64(m.grossMoney.Amount()) * float64(m.vat.denominator) / float64(m.vat.denominator+m.vat.rate)
		m.netMoney = money.New(
			int64(math.RoundToEven(amount)),
			m.grossMoney.Currency().Code,
		)
	} else {
		amount := float64(m.netMoney.Amount()) * float64(m.vat.denominator+m.vat.rate) / float64(m.vat.denominator)
		m.grossMoney = money.New(
			int64(math.RoundToEven(amount)),
			m.netMoney.Currency().Code,
		)
	}
	return m
}

func (m Money) DisplayNetOrGross(gross bool, opts ...DisplayOption) string {
	if gross {
		return m.DisplayGross(opts...)
	}
	return m.DisplayNet(opts...)
}

func (m Money) Display(opts ...DisplayOption) string {
	return m.DisplayNetOrGross(m.LeadingValueIsGross, opts...)
}

func (m Money) DisplayGross(opts ...DisplayOption) string {
	moneyStr := m.grossMoney.Display()
	do := &displayOptions{}
	for _, opt := range opts {
		opt(do)
	}
	if do.DisplayVatSuffix {
		moneyStr = fmt.Sprintf(do.GrossVatSuffixFormat, moneyStr, float64(m.vat.rate)/float64(m.vat.denominator)*100.0)
	}
	return moneyStr
}

func (m Money) DisplayNet(opts ...DisplayOption) string {
	moneyStr := m.netMoney.Display()
	do := &displayOptions{}
	for _, opt := range opts {
		opt(do)
	}
	if do.DisplayVatSuffix {
		moneyStr = fmt.Sprintf(do.NetVatSuffixFormat, moneyStr, float64(m.vat.rate)/float64(m.vat.denominator)*100.0)
	}
	return moneyStr
}

func (m *Money) calculateWithMoney() *money.Money {
	if m.LeadingValueIsGross {
		return m.grossMoney
	}
	return m.netMoney
}

func (m *Money) newWithAmount(om *money.Money) *Money {
	var rm *Money
	if m.LeadingValueIsGross {
		rm = &Money{
			grossMoney:          om,
			vat:                 m.vat,
			LeadingValueIsGross: true,
		}
	} else {
		rm = &Money{
			netMoney:            om,
			vat:                 m.vat,
			LeadingValueIsGross: false,
		}
	}
	return rm.updateOtherValue()
}

func (m *Money) netOrGross(gross bool) *money.Money {
	if gross {
		return m.grossMoney
	}
	return m.netMoney
}
