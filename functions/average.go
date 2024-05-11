package functions

import "strings"

func Average(inputFile []byte) int{
	inputStr := strings.Split(string(inputFile), "\n")
	count := 0
	sum := 0
	for i := 0; i < len(inputStr); i++{
		count += 1
		sum += StrToInt(inputStr[i])

	}
	average := sum/count
	return average
}
