package adventutils

import (
	"errors"
	"os"
)

func findTwoTermsFor(sum int, terms []int) (a, b int, err error) {
	for indexA, a := range terms {
		termsSlice := terms[indexA+1 : len(terms)]
		for _, b := range termsSlice {
			if a+b == sum {
				return a, b, nil
			}
		}
	}
	return 0, 0, errors.New("Could not find numbers")
}

func findThreeTermsFor(sum int, terms []int) (a, b, c int, err error) {
	for indexA, a := range terms {
		termsSliceA := terms[indexA+1 : len(terms)]
		for indexB, b := range termsSliceA {
			termsSliceB := termsSliceA[indexB+1 : len(termsSliceA)]
			for _, c := range termsSliceB {
				if a+b+c == sum {
					return a, b, c, nil
				}
			}
		}
	}

	return 0, 0, 0, errors.New("Could not find numbers")
}

func OneA(invoices []int) int {
	a, b, err := findTwoTermsFor(2020, invoices)
	if errors.Is(err, os.ErrExist) {
		return 0
	}

	return a * b
}

func OneB(invoices []int) int {
	a, b, c, err := findThreeTermsFor(2020, invoices)
	if errors.Is(err, os.ErrExist) {
		return 0
	}

	return a * b * c
}
