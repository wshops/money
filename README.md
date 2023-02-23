# WSHOPS / MONEY

a money utils used in all WSHOPS projects with a good performance and a good API.

## Installation

```bash
go get github.com/wshops/money
```

## Usage

```go
package main

import "github.com/wshops/money"

func main() {
	m, err := money.NewFromString("1.23", money.CNY)
	// handle err...

	m.AddStr("1.23").AddDollarInt(10).ToString()
}
```

see tests and benchmark for more details.

## Benchmark

```bash
go test -bench=. -benchmem -benchtime=3s -cpu=10
```

result:

```bash
Benchmarking...
goos: darwin
goarch: arm64
pkg: github.com/wshops/money
BenchmarkNewFromString-10               26983059               115.3 ns/op            56 B/op          2 allocs/op
BenchmarkMustFromString-10              30604957               116.1 ns/op            56 B/op          2 allocs/op
BenchmarkAdd-10                         1000000000               0.7636 ns/op          0 B/op          0 allocs/op
BenchmarkAddCentsInt-10                 1000000000               0.3193 ns/op          0 B/op          0 allocs/op
BenchmarkAddDollarInt-10                1000000000               0.3208 ns/op          0 B/op          0 allocs/op
BenchmarkSubStr-10                      59271128                59.38 ns/op           32 B/op          1 allocs/op
BenchmarkSub-10                         1000000000               0.7603 ns/op          0 B/op          0 allocs/op
BenchmarkSubCentsInt-10                 1000000000               0.3217 ns/op          0 B/op          0 allocs/op
BenchmarkSubDollarInt-10                1000000000               0.3176 ns/op          0 B/op          0 allocs/op
BenchmarkToString-10                    37875367                94.00 ns/op           35 B/op          1 allocs/op
BenchmarkToStringCurrency-10            22697878               162.3 ns/op            55 B/op          2 allocs/op
BenchmarkToCentsInt-10                  288610015               13.77 ns/op           29 B/op          0 allocs/op
BenchmarkIsValidStrAmount-10            72290002                49.36 ns/op            0 B/op          0 allocs/op
BenchmarkParseStrToCentsInt-10          14065726               245.6 ns/op           102 B/op          5 allocs/op
PASS
ok      github.com/wshops/money 36.099s
```