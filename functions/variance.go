package functions

import "strings"

func Variance(inputFile []byte) int{
	mean := Average(inputFile)
	
	inputStr := strings.Split(string(inputFile), "\n")
	sum := 0
	count := 0

	for i := 0; i < len(inputStr); i++{
		count++
		diff := StrToInt(inputStr[i]) - mean
		diff2 := (diff*diff)
		sum += diff2
	}
	result := sum/count
	return result
}