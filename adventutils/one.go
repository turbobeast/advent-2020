package adventutils

func OneA(invoices []int) int {
	for indexA, a := range invoices {
		matches := invoices[indexA+1 : len(invoices)]
		for _, b := range matches {
			if a+b == 2020 {
				return a * b
			}
		}
	}
	return 0
}

func OneB(invoices []int) int {
	for indexA, a := range invoices {
		firstMatches := invoices[indexA+1 : len(invoices)]
		for indexB, b := range firstMatches {
			secondMatches := firstMatches[indexB+1 : len(firstMatches)]
			for _, c := range secondMatches {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}

	return 0
}
