package main

import (
	"fmt"
	"advent-2020/adventutils"
)

func main () {
	invoices := adventutils.ParseInput("input/one.txt")
	prod := adventutils.OneB(invoices)
	fmt.Println(prod)
}
