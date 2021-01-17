package main

import (
	"advent-2020/adventutils"
	"fmt"
)

func main() {
	invoices := adventutils.ParseInput("input/one.txt")
	prod := adventutils.OneB(invoices)
	fmt.Println(prod)
}
