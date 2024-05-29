package functions

import (
	"sort"
)

func Median(inputFile []float64) float64 {
	sort.Float64s(inputFile) // sort the slice in ascending order

	var num1, num2, median float64
	n := len(inputFile)

	if n%2 == 0 { //calculate median differently for even and odd length
		num1 = inputFile[(n/2)-1]
		num2 = inputFile[(n / 2)]
		median = (num1 + num2) / 2
	} else {
		median = inputFile[(n / 2)]
	}

	return median
}
