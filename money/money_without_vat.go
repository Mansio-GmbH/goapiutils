package money

import (
	"math"

	"github.com/Rhymond/go-money"
)

type MoneyWithoutVat struct {
	money money.Money
}

func NewWithoutVat(amount int64, currency string) *MoneyWithoutVat {
	return &MoneyWithoutVat{money: *money.New(amount, currency)}
}

func NewWithoutVatFromFloat(amount float64, currency string) *MoneyWithoutVat {
	return &MoneyWithoutVat{money: *money.NewFromFloat(amount, currency)}
}

func NewWithoutVatFromMoney(money *money.Money) *MoneyWithoutVat {
	return &MoneyWithoutVat{money: *money}
}

func (m *MoneyWithoutVat) Add(om *MoneyWithoutVat) (*MoneyWithoutVat, error) {
	if om == nil {
		return m, nil
	}
	rm, err := m.money.Add(&om.money)
	if err != nil {
		return nil, err
	}
	return &MoneyWithoutVat{money: *rm}, nil
}

func (m *MoneyWithoutVat) Multiply(n int64) *MoneyWithoutVat {
	return &MoneyWithoutVat{money: *m.money.Multiply(n)}
}

func (m *MoneyWithoutVat) MultiplyByFloat(x float64) *MoneyWithoutVat {
	newAmount := int64(math.Round(float64(m.Amount()) * x))
	return NewWithoutVat(newAmount, m.CurrencyCode())
}

func (m *MoneyWithoutVat) Percentage(perc float64) *MoneyWithoutVat {
	return m.MultiplyByFloat(perc / 100.0)
}

func (m *MoneyWithoutVat) Amount() int64 {
	return m.money.Amount()
}

func (m *MoneyWithoutVat) CurrencyCode() string {
	return m.money.Currency().Code
}

func (m *MoneyWithoutVat) Equals(om *MoneyWithoutVat) (bool, error) {
	return m.money.Equals(&om.money)
}

func (m *MoneyWithoutVat) IsZero() bool {
	return m.money.IsZero()
}

func (m *MoneyWithoutVat) IsPositive() bool {
	return m.money.IsPositive()
}

func (m *MoneyWithoutVat) IsNegative() bool {
	return m.money.IsNegative()
}

func (m *MoneyWithoutVat) LessThan(om *MoneyWithoutVat) (bool, error) {
	return m.money.LessThan(&om.money)
}

func (m *MoneyWithoutVat) LessThanOrEqual(om *MoneyWithoutVat) (bool, error) {
	return m.money.LessThanOrEqual(&om.money)
}

func (m *MoneyWithoutVat) GreaterThan(om *MoneyWithoutVat) (bool, error) {
	return m.money.GreaterThan(&om.money)
}

func (m *MoneyWithoutVat) GreaterThanOrEqual(om *MoneyWithoutVat) (bool, error) {
	return m.money.GreaterThanOrEqual(&om.money)
}

func (m *MoneyWithoutVat) Negate() *MoneyWithoutVat {
	return NewWithoutVat(-m.Amount(), m.CurrencyCode())
}

func (m *MoneyWithoutVat) Negative() *MoneyWithoutVat {
	return NewWithoutVat(m.money.Negative().Amount(), m.CurrencyCode())
}

func (m *MoneyWithoutVat) Absolute() *MoneyWithoutVat {
	return NewWithoutVat(m.money.Absolute().Amount(), m.CurrencyCode())
}

func (m *MoneyWithoutVat) Display() string {
	return m.money.Display()
}
