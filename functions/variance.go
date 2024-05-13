package functions

import "strings"

func Variance(inputFile []byte) int{
	mean := Average(inputFile)
	
	input := strings.Split(string(inputFile), "\n")
	inputStr := RemoveInvalid(input)

	var sum, count int

	for i := 0; i < len(inputStr); i++{
		count++
		diff := StrToInt(inputStr[i]) - mean
		diff2 := (diff*diff)
		sum += diff2
	}
	result := sum/count
	return result
}