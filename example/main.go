package main

import (
	"fmt"
	"github.com/wshops/money"
)

func main() {
	m, err := money.NewFromString("1.23", money.CNY)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m.AddStr("1.23").AddDollarInt(10).ToString())
}
