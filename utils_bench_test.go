package money

import (
	"testing"
)

func BenchmarkIsValidStrAmount(b *testing.B) {
	inputs := []string{
		"123.456",
		"1000",
		"1.234.567",
		"",
		" ",
		"123 456",
		"1234",
		"1.23.4",
		"123.4.5",
	}

	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			if IsValidStrAmount(input) {
				// do nothing
			}
		}
	}
}
func BenchmarkParseStrToCentsInt(b *testing.B) {
	inputs := []string{
		"100",
		"123.45",
		".3",
		"0.3",
		"123456789.01",
	}

	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			parseStrToCentsInt(input)
		}
	}
}
