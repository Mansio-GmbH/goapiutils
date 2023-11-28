package money

import (
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
