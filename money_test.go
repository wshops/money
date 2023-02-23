package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFromString(t *testing.T) {
	// Test valid input with default currency
	m, err := NewFromString("10.00")
	assert.Nil(t, err, "NewFromString should not return an error for valid input")
	assert.Equal(t, int64(1000), m.amountCents, "NewFromString should create a Money instance with the correct amountCents")
	assert.Equal(t, USD, m.currency, "NewFromString should create a Money instance with the default currency (USD)")

	// Test valid input with specified currency
	m, err = NewFromString("10.00", EUR)
	assert.Nil(t, err, "NewFromString should not return an error for valid input")
	assert.Equal(t, int64(1000), m.amountCents, "NewFromString should create a Money instance with the correct amountCents")
	assert.Equal(t, EUR, m.currency, "NewFromString should create a Money instance with the specified currency")

	// Test invalid input
	m, err = NewFromString("invalid")
	assert.NotNil(t, err, "NewFromString should return an error for invalid input")
	assert.Nil(t, m, "NewFromString should not return a Money instance for invalid input")

	// Test invalid currency
	m, err = NewFromString("10.00", "XYZ")
	assert.NotNil(t, err, "NewFromString should return an error for an invalid currency")
	assert.Nil(t, m, "NewFromString should not return a Money instance for an invalid currency")
}

func TestMustFromString(t *testing.T) {
	// Test valid input with default currency
	m := MustFromString("10.00")
	assert.Equal(t, int64(1000), m.amountCents, "MustFromString should create a Money instance with the correct amountCents")
	assert.Equal(t, USD, m.currency, "MustFromString should create a Money instance with the default currency (USD)")

	// Test valid input with specified currency
	m = MustFromString("10.00", EUR)
	assert.Equal(t, int64(1000), m.amountCents, "MustFromString should create a Money instance with the correct amountCents")
	assert.Equal(t, EUR, m.currency, "MustFromString should create a Money instance with the specified currency")

	// Test invalid input (should panic)
	assert.Panics(t, func() {
		MustFromString("invalid")
	}, "MustFromString should panic for invalid input")

	// Test invalid currency (should panic)
	assert.Panics(t, func() {
		MustFromString("10.00", "XYZ")
	}, "MustFromString should panic for an invalid currency")
}

func TestAdd(t *testing.T) {
	m1 := MustFromString("10.00")
	m2 := MustFromString("5.00")
	expected := MustFromString("15.00")
	result := m1.Add(m2)
	assert.Equal(t, expected, result, "Add result should match expected")
}

func TestAddCentsInt(t *testing.T) {
	m1 := MustFromString("10.00")
	c := 500
	expected := MustFromString("15.00")
	result := m1.AddCentsInt(c)
	assert.Equal(t, expected, result, "AddCentsInt result should match expected")
}

func TestAddDollarInt(t *testing.T) {
	m1 := MustFromString("10.00")
	d := 5
	expected := MustFromString("15.00")
	result := m1.AddDollarInt(d)
	assert.Equal(t, expected, result, "AddDollarInt result should match expected")
}

func TestSubStr(t *testing.T) {
	m1 := MustFromString("10.00")
	expected := MustFromString("5.00")
	result := m1.SubStr("5.00")
	assert.Equal(t, expected, result, "SubStr result should match expected")
}

func TestSub(t *testing.T) {
	m1 := MustFromString("10.00")
	m2 := MustFromString("5.00")
	expected := MustFromString("5.00")
	result := m1.Sub(m2)
	assert.Equal(t, expected, result, "Sub result should match expected")
}

func TestSubCentsInt(t *testing.T) {
	m1 := MustFromString("10.00")
	c := 500
	expected := MustFromString("5.00")
	result := m1.SubCentsInt(c)
	assert.Equal(t, expected, result, "SubCentsInt result should match expected")
}

func TestSubDollarInt(t *testing.T) {
	m1 := MustFromString("10.00")
	d := 5
	expected := MustFromString("5.00")
	result := m1.SubDollarInt(d)
	assert.Equal(t, expected, result, "SubDollarInt result should match expected")
}
