package money

import (
	"errors"
	"fmt"
)

type Money struct {
	amountCents int64
	currency    Currency
}

const (
	ErrInvalidAmountStr = "invalid amount string"
	ErrTypeConversion   = "type conversion error"

	ErrInvalidCurrency = "invalid currency"
)

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	CNY Currency = "CNY"
)

// NewFromString
// @Description: NewFromString creates a new Money instance from a string.
// @Param amount
// @Param currency
func NewFromString(amount string, currency ...Currency) (*Money, error) {
	// validate input
	if !isValidStrAmount(amount) {
		return nil, errors.New(ErrInvalidAmountStr)
	}
	// get instance
	m := acquireMoney()
	// currency
	if len(currency) > 0 {
		if !isValidCurrency(string(currency[0])) {
			return nil, errors.New(ErrInvalidCurrency)
		}
		m.currency = currency[0]
	} else {
		m.currency = USD
	}
	var err error
	m.amountCents, err = parseStrToCentsInt(amount)
	if err != nil {
		return nil, err
	}
	// return
	return m, nil
}

// MustFromString
// @Description: MustFromString is a helper function that wraps NewFromString and panics if an error is returned.
// @Param amount
// @Param currency
func MustFromString(amount string, currency ...Currency) *Money {
	m, err := NewFromString(amount, currency...)
	if err != nil {
		panic(err)
	}
	return m
}

// AddStr
// @Description: AddStr adds a string amount to the Money instance.
// @Param str
func (m *Money) AddStr(str string) *Money {
	// validate input
	if !isValidStrAmount(str) {
		panic(ErrInvalidAmountStr)
	}
	// parse
	amount, err := parseStrToCentsInt(str)
	if err != nil {
		panic(err)
	}
	// add
	m.amountCents += amount
	// return
	return m
}

// Add
// @Description: Add adds a Money instance to the Money instance.
// @Param money
func (m *Money) Add(money *Money) *Money {
	// add
	m.amountCents += money.amountCents
	// return
	return m
}

// AddCentsInt
// @Description: AddCentsInt adds a cents int to the Money instance.
// @Param c
func (m *Money) AddCentsInt(c int) *Money {
	// add
	m.amountCents += int64(c)
	// return
	return m
}

// AddDollarInt
// @Description: AddDollarInt adds a dollar int to the Money instance.
// @Param d
func (m *Money) AddDollarInt(d int) *Money {
	// add
	m.amountCents += int64(d * 100)
	// return
	return m
}

// SubStr
// @Description: SubStr subtracts a string amount from the Money instance.
// @Param str
func (m *Money) SubStr(str string) *Money {
	// validate input
	if !isValidStrAmount(str) {
		panic(ErrInvalidAmountStr)
	}
	// parse
	amount, err := parseStrToCentsInt(str)
	if err != nil {
		panic(err)
	}
	// sub
	m.amountCents -= amount
	// return
	return m
}

// Sub
// @Description: Sub subtracts a Money instance from the Money instance.
// @Param money
func (m *Money) Sub(money *Money) *Money {
	// sub
	m.amountCents -= money.amountCents
	// return
	return m
}

// SubCentsInt
// @Description: SubCentsInt subtracts a cents int from the Money instance.
// @Param c
func (m *Money) SubCentsInt(c int) *Money {
	// sub
	m.amountCents -= int64(c)
	// return
	return m
}

// SubDollarInt
// @Description: SubDollarInt subtracts a dollar int from the Money instance.
// @Param d
func (m *Money) SubDollarInt(d int) *Money {
	// sub
	m.amountCents -= int64(d * 100)
	// return
	return m
}

func (m *Money) ToString() string {
	defer releaseMoney(m)
	return fmt.Sprintf("%.2f", float64(m.amountCents)/100)
}

func (m *Money) ToStringCurrency() string {
	defer releaseMoney(m)
	return fmt.Sprintf("%s %.2f", m.currency, float64(m.amountCents)/100)
}
func (m *Money) ToCentsInt() int64 {
	defer releaseMoney(m)
	result := m.amountCents
	return result
}
