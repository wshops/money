package money

import (
	"errors"
	"strconv"
	"strings"
)

type Money *money
type money struct {
	AmountCents int64
	Currency    Currency
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
		m.Currency = currency[0]
	} else {
		m.Currency = USD
	}
	var decPart int64
	var dolPart int64
	var err error
	// in case of no decimal part
	if !strings.Contains(amount, ".") {
		// convert dollar part to cents int
		dolPart, err = strconv.ParseInt(amount, 10, 64)
		if err != nil {
			return nil, errors.New(ErrTypeConversion)
		}
	} else if strings.HasPrefix(amount, ".") {
		// append 0 if only one digit for decimal part
		if len(amount[1:]) == 1 {
			amount += "0"
		}
		// convert decimal part to cents int
		decPart, err = strconv.ParseInt(amount[1:], 10, 64)
		if err != nil {
			return nil, errors.New(ErrTypeConversion)
		}
	} else {
		// in case of dollar and decimal part
		splitAmount := strings.Split(amount, ".")
		// append 0 if only one digit for decimal part
		if len(splitAmount[1]) == 1 {
			splitAmount[1] += "0"
		}
		// convert decimal part to cents int
		decPart, err = strconv.ParseInt(splitAmount[1], 10, 64)
		if err != nil {
			return nil, errors.New(ErrTypeConversion)
		}
		// convert dollar part to cents int
		dolPart, err = strconv.ParseInt(splitAmount[0], 10, 64)
		if err != nil {
			return nil, errors.New(ErrTypeConversion)
		}
	}
	// calculate amount in cents
	m.AmountCents = dolPart*100 + decPart
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
