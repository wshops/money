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
			if isValidStrAmount(input) {
				// do nothing
			}
		}
	}
}

func BenchmarkNewFromStringDefaultCurrency(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m, _ := NewFromString("10.50")
		releaseMoney(m)
	}
}

func BenchmarkNewFromStringWithCurrency(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m, _ := NewFromString("100", EUR)
		releaseMoney(m)
	}
}

func BenchmarkNewFromStringNoDecimal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m, _ := NewFromString("10")
		releaseMoney(m)
	}
}

func BenchmarkMustFromStringDefaultCurrency(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MustFromString("10.50")
	}
}

func BenchmarkMustFromStringWithCurrency(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MustFromString("100", EUR)
	}
}

func BenchmarkMustFromStringNoDecimal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MustFromString("10")
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
