package money

import (
	"errors"
	"testing"
)

func TestIsValidAmountStr(t *testing.T) {
	validInputs := []string{
		"123.45",
		"1212",
		"100",
		"1",
		"0",
		"0.0",
		"12.0",
		"0.00",
	}
	for _, input := range validInputs {
		allowed := isValidStrAmount(input)
		if !allowed {
			t.Errorf("invalid input: %s", input)
			t.Failed()
		}
	}

	invalidInputs := []string{
		"123.456.789.123",
		"",
		" ",
		"asdad",
		"as.",
		"a.s",
		".asd",
		"98.9999",
	}
	for _, input := range invalidInputs {
		allowed := isValidStrAmount(input)
		if allowed {
			t.Errorf("invalid input identify error: %s", input)
			t.Failed()
		}
	}
}

func TestNewFromString(t *testing.T) {
	// Test case 1
	m, err := NewFromString("10.50")
	if err != nil {
		t.Errorf("Test case 1 failed with error: %v", err)
	}
	if m.amountCents != 1050 || m.currency != USD {
		t.Errorf("Test case 1 failed with unexpected output: %v", m)
	}

	// Test case 2
	m, err = NewFromString("100", EUR)
	if err != nil {
		t.Errorf("Test case 2 failed with error: %v", err)
	}
	if m.amountCents != 10000 || m.currency != EUR {
		t.Errorf("Test case 2 failed with unexpected output: %v", m)
	}

	// Test case 3
	_, err = NewFromString("10.5.5")
	if err == nil || err.Error() != ErrInvalidAmountStr {
		t.Errorf("Test case 3 failed with unexpected error: %v", err)
	}

	// Test case 4
	_, err = NewFromString("abc")
	if err == nil || err.Error() != ErrInvalidAmountStr {
		t.Errorf("Test case 4 failed with unexpected error: %v", err)
	}

	// Test case 5
	m, err = NewFromString("10.5")
	if err != nil {
		t.Errorf("Test case 5 failed with error: %v", err)
	}
	if m.amountCents != 1050 || m.currency != USD {
		t.Errorf("Test case 5 failed with unexpected output: %v", m)
	}

	// Test case 6
	m, err = NewFromString("10.5", CNY)
	if err != nil {
		t.Errorf("Test case 6 failed with error: %v", err)
	}
	if m.amountCents != 1050 || m.currency != CNY {
		t.Errorf("Test case 6 failed with unexpected output: %v", m)
	}

	// Test case 7
	m, err = NewFromString("10.51")
	if err != nil {
		t.Errorf("Test case 7 failed with error: %v", err)
	}
	if m.amountCents != 1051 || m.currency != USD {
		t.Errorf("Test case 7 failed with unexpected output: %v", m)
	}

	// Test case 8
	m, err = NewFromString("10.5")
	if err != nil {
		t.Errorf("Test case 8 failed with error: %v", err)
	}
	if m.amountCents != 1050 || m.currency != USD {
		t.Errorf("Test case 8 failed with unexpected output: %v", m)
	}

	// Test case 9
	m, err = NewFromString(".5")
	if err != nil {
		t.Errorf("Test case 9 failed with error: %v", err)
	}
	if m.amountCents != 50 || m.currency != USD {
		t.Errorf("Test case 9 failed with unexpected output: %v", m)
	}

	// Test case 10
	m, err = NewFromString(".05")
	if err != nil {
		t.Errorf("Test case 10 failed with error: %v", err)
	}
	if m.amountCents != 5 || m.currency != USD {
		t.Errorf("Test case 10 failed with unexpected output: %v", m)
	}

	// Test case 11
	m, err = NewFromString("0.05")
	if err != nil {
		t.Errorf("Test case 11 failed with error: %v", err)
	}
	if m.amountCents != 5 || m.currency != USD {
		t.Errorf("Test case 11 failed with unexpected output: %v", m)
	}

	// Test case 12
	m, err = NewFromString("0.50")
	if err != nil {
		t.Errorf("Test case 12 failed with error: %v", err)
	}
	if m.amountCents != 50 || m.currency != USD {
		t.Errorf("Test case 12 failed with unexpected output: %v", m)
	}
}

func TestMustFromString(t *testing.T) {
	// Test case 1
	m := MustFromString("10.50")
	if m.amountCents != 1050 || m.currency != USD {
		t.Errorf("Test case 1 failed with unexpected output: %v", m)
	}

	// Test case 2
	m = MustFromString("100", EUR)
	if m.amountCents != 10000 || m.currency != EUR {
		t.Errorf("Test case 2 failed with unexpected output: %v", m)
	}

	// Test case 3
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test case 3 failed with no panic")
		}
	}()
	MustFromString("10.5.5")
}

func TestParseStrToCentsInt(t *testing.T) {
	testCases := []struct {
		input string
		want  int64
		err   error
	}{
		{
			input: "100",
			want:  10000,
			err:   nil,
		},
		{
			input: "123.45",
			want:  12345,
			err:   nil,
		},
		{
			input: ".3",
			want:  30,
			err:   nil,
		},
		{
			input: "0.3",
			want:  30,
			err:   nil,
		},
		{
			input: "abc",
			want:  0,
			err:   errors.New(ErrTypeConversion),
		},
	}

	for _, tc := range testCases {
		got, err := parseStrToCentsInt(tc.input)
		if got != tc.want {
			t.Errorf("parseStrToCentsInt(%q) = %d, want %d", tc.input, got, tc.want)
		}
		if (err == nil && tc.err != nil) || (err != nil && tc.err == nil) || (err != nil && tc.err != nil && err.Error() != tc.err.Error()) {
			t.Errorf("parseStrToCentsInt(%q) error = %v, want %v", tc.input, err, tc.err)
		}
	}
}
