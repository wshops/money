package money

import (
	"errors"
	"fmt"
)

type Money *money
type money struct {
	amountCents int64
	currency    Currency
}

// NewFromString
// @Description: NewFromString creates a new Money instance from a string.
// @Param amount
// @Param currency
func NewFromString(amount string, currency ...Currency) (Money, error) {
	// validate input
	if !isValidStrAmount(amount) {
		return nil, errors.New(ErrInvalidAmountStr)
	}
	// get instance
	m := acquireMoney()
	// currency
	if len(currency) > 0 {
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
func MustFromString(amount string, currency ...Currency) Money {
	m, err := NewFromString(amount, currency...)
	if err != nil {
		panic(err)
	}
	return m
}

// AddStr
// @Description: AddStr adds a string amount to the Money instance.
// @Param str
func (m *money) AddStr(str string) Money {
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
func (m *money) Add(money Money) Money {
	// add
	m.amountCents += money.amountCents
	// return
	return m
}

// AddCentsInt
// @Description: AddCentsInt adds a cents int to the Money instance.
// @Param c
func (m *money) AddCentsInt(c int) Money {
	// add
	m.amountCents += int64(c)
	// return
	return m
}

// AddDollarInt
// @Description: AddDollarInt adds a dollar int to the Money instance.
// @Param d
func (m *money) AddDollarInt(d int) Money {
	// add
	m.amountCents += int64(d * 100)
	// return
	return m
}

// SubStr
// @Description: SubStr subtracts a string amount from the Money instance.
// @Param str
func (m *money) SubStr(str string) Money {
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
func (m *money) Sub(money Money) Money {
	// sub
	m.amountCents -= money.amountCents
	// return
	return m
}

// SubCentsInt
// @Description: SubCentsInt subtracts a cents int from the Money instance.
// @Param c
func (m *money) SubCentsInt(c int) Money {
	// sub
	m.amountCents -= int64(c)
	// return
	return m
}

// SubDollarInt
// @Description: SubDollarInt subtracts a dollar int from the Money instance.
// @Param d
func (m *money) SubDollarInt(d int) Money {
	// sub
	m.amountCents -= int64(d * 100)
	// return
	return m
}

// MulStr
// @Description: MulStr multiplies the Money instance by a string amount.
// @Param str
func (m *money) MulStr(str string) Money {
	// validate input
	if !isValidStrAmount(str) {
		panic(ErrInvalidAmountStr)
	}
	// parse
	amount, err := parseStrToCentsInt(str)
	if err != nil {
		panic(err)
	}
	// mul
	m.amountCents *= amount
	// return
	return m
}

// Mul
// @Description: Mul multiplies the Money instance by a Money instance.
// @Param money
func (m *money) Mul(money Money) Money {
	// mul
	m.amountCents *= money.amountCents
	// return
	return m
}

// MulCentsInt
// @Description: MulCentsInt multiplies the Money instance by a cents int.
// @Param c
func (m *money) MulCentsInt(c int) Money {
	// mul
	m.amountCents *= int64(c)
	// return
	return m
}

// MulDollarInt
// @Description: MulDollarInt multiplies the Money instance by a dollar int.
// @Param d
func (m *money) MulDollarInt(d int) Money {
	// mul
	m.amountCents *= int64(d * 100)
	// return
	return m
}

// DivStr
// @Description: DivStr divides the Money instance by a string amount.
// @Param str
func (m *money) DivStr(str string) Money {
	// validate input
	if !isValidStrAmount(str) {
		panic(ErrInvalidAmountStr)
	}
	// parse
	amount, err := parseStrToCentsInt(str)
	if err != nil {
		panic(err)
	}
	// div
	m.amountCents /= amount
	// return
	return m
}

// Div
// @Description: Div divides the Money instance by a Money instance.
// @Param money
func (m *money) Div(money Money) Money {
	// div
	m.amountCents /= money.amountCents
	// return
	return m
}

// DivCentsInt
// @Description: DivCentsInt divides the Money instance by a cents int.
// @Param c
func (m *money) DivCentsInt(c int) Money {
	// div
	m.amountCents /= int64(c)
	// return
	return m
}

// DivDollarInt
// @Description: DivDollarInt divides the Money instance by a dollar int.
// @Param d
func (m *money) DivDollarInt(d int) Money {
	// div
	m.amountCents /= int64(d * 100)
	// return
	return m
}

func (m *money) ToString() string {
	defer releaseMoney(m)
	return fmt.Sprintf("%.2f", float64(m.amountCents)/100)
}

func (m *money) ToCentsInt() int64 {
	defer releaseMoney(m)
	return m.amountCents
}
