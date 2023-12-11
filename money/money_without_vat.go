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

func (m *MoneyWithoutVat) Display() string {
	return m.money.Display()
}
