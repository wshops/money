package money

import "testing"

func BenchmarkNewFromString(b *testing.B) {
	// Reset the benchmark timer
	b.ResetTimer()

	// Create Money instances from a string repeatedly
	for i := 0; i < b.N; i++ {
		_, err := NewFromString("10.00")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMustFromString(b *testing.B) {
	// Reset the benchmark timer
	b.ResetTimer()

	// Create Money instances from a string repeatedly
	for i := 0; i < b.N; i++ {
		MustFromString("10.00")
	}
}

func BenchmarkAdd(b *testing.B) {
	// Create two Money instances to add together
	m1 := MustFromString("10.00")
	m2 := MustFromString("5.00")

	// Reset the benchmark timer
	b.ResetTimer()

	// Add the two Money instances together repeatedly
	for i := 0; i < b.N; i++ {
		m1.Add(m2)
	}
}

func BenchmarkAddCentsInt(b *testing.B) {
	// Create a Money instance to add to
	m := MustFromString("10.00")

	// Reset the benchmark timer
	b.ResetTimer()

	// Add a cents int repeatedly
	for i := 0; i < b.N; i++ {
		m.AddCentsInt(500)
	}
}

func BenchmarkAddDollarInt(b *testing.B) {
	// Create a Money instance to add to
	m := MustFromString("10.00")

	// Reset the benchmark timer
	b.ResetTimer()

	// Add a dollar int repeatedly
	for i := 0; i < b.N; i++ {
		m.AddDollarInt(1)
	}
}

func BenchmarkSubStr(b *testing.B) {
	// Create a Money instance to subtract from
	m := MustFromString("10.00")

	// Reset the benchmark timer
	b.ResetTimer()

	// Subtract a string amount repeatedly
	for i := 0; i < b.N; i++ {
		m.SubStr("5.00")
	}
}

func BenchmarkSub(b *testing.B) {
	// Create two Money instances to subtract
	m1 := MustFromString("10.00")
	m2 := MustFromString("5.00")

	// Reset the benchmark timer
	b.ResetTimer()

	// Subtract one Money instance from the other repeatedly
	for i := 0; i < b.N; i++ {
		m1.Sub(m2)
	}
}

func BenchmarkSubCentsInt(b *testing.B) {
	// Create a Money instance to subtract from
	m := MustFromString("10.00")

	// Reset the benchmark timer
	b.ResetTimer()

	// Subtract a cents int repeatedly
	for i := 0; i < b.N; i++ {
		m.SubCentsInt(500)
	}
}

func BenchmarkSubDollarInt(b *testing.B) {
	// Create a Money instance to subtract from
	m := MustFromString("10.00")

	// Reset the benchmark timer
	b.ResetTimer()

	// Subtract a dollar int repeatedly
	for i := 0; i < b.N; i++ {
		m.SubDollarInt(1)
	}
}

func BenchmarkToString(b *testing.B) {
	// Create a Money instance to convert to a string
	m := MustFromString("10.00")

	// Reset the benchmark timer
	b.ResetTimer()

	// Convert the Money instance to a string repeatedly
	for i := 0; i < b.N; i++ {
		m.ToString()
	}
}

func BenchmarkToStringCurrency(b *testing.B) {
	// Create a Money instance to convert to a string with currency
	m := MustFromString("10.00", EUR)

	// Reset the benchmark timer
	b.ResetTimer()

	// Convert the Money instance to a string with currency repeatedly
	for i := 0; i < b.N; i++ {
		m.ToStringCurrency()
	}
}

func BenchmarkToCentsInt(b *testing.B) {
	// Create a Money instance to get the amountCents from
	m := MustFromString("10.00")

	// Reset the benchmark timer
	b.ResetTimer()

	// Get the amountCents from the Money instance repeatedly
	for i := 0; i < b.N; i++ {
		m.ToCentsInt()
	}
}
