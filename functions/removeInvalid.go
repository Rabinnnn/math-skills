package functions

import(
	"strconv"
	"fmt"
)

func RemoveInvalid(input []string) []string{
	var output []string

	for _, num := range input{
		floatingNum, err := strconv.ParseFloat(num, 64)
		if err != nil{
			fmt.Println("Error:", err)
		}

		if floatingNum >= 0.0 && num != "" || num == "."{
			output = append(output, num)
		} 
		
	}
	return output
}
