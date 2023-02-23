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
