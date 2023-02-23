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
BenchmarkNewFromString-10               27203233               116.3 ns/op            56 B/op          2 allocs/op
BenchmarkMustFromString-10              31251073               116.3 ns/op            56 B/op          2 allocs/op
BenchmarkAdd-10                         1000000000               0.7662 ns/op          0 B/op          0 allocs/op
BenchmarkAddCentsInt-10                 1000000000               0.3175 ns/op          0 B/op          0 allocs/op
BenchmarkAddDollarInt-10                1000000000               0.3213 ns/op          0 B/op          0 allocs/op
BenchmarkSubStr-10                      59678316                59.57 ns/op           32 B/op          1 allocs/op
BenchmarkSub-10                         1000000000               0.7605 ns/op          0 B/op          0 allocs/op
BenchmarkSubCentsInt-10                 1000000000               0.3208 ns/op          0 B/op          0 allocs/op
BenchmarkSubDollarInt-10                1000000000               0.3209 ns/op          0 B/op          0 allocs/op
BenchmarkToString-10                    38568621                93.96 ns/op           35 B/op          1 allocs/op
BenchmarkToStringCurrency-10            22570384               162.7 ns/op            54 B/op          2 allocs/op
BenchmarkToCentsInt-10                  289517320               13.41 ns/op           29 B/op          0 allocs/op
BenchmarkMul-10                         1000000000               0.3196 ns/op          0 B/op          0 allocs/op
BenchmarkDiv-10                         1000000000               1.400 ns/op           0 B/op          0 allocs/op
BenchmarkIsValidStrAmount-10            71433060                49.71 ns/op            0 B/op          0 allocs/op
BenchmarkParseStrToCentsInt-10          14425176               245.9 ns/op           102 B/op          5 allocs/op
PASS
ok      github.com/wshops/money 38.484s
```