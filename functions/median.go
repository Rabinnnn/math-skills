package functions

import(
	"strings"
)


func Median(inputFile []byte) int{
	inputStr := strings.Split(string(inputFile), "\n")
	
	CustomSort(inputStr) // sort the slice in ascending order

	var num1, num2, median int
	n := len(inputStr)
	
	for i := 0; i < n; i++{
		if len(inputStr)%2 == 0{ //check if the length is an even integer.
			num1 = StrToInt(inputStr[(n/2)-1])
			num2 = StrToInt(inputStr[(n/2)])
			median = (num1+num2)/2
		} else{
			median = StrToInt(inputStr[(n/2)])
		}
			
	}
	return median
}
