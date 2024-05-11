package functions

import "fmt"

func StrToInt(input string) int{
	result := 0
	sign := 1

	for i, char := range input {
		// Check for the sign (+ or -) at the beginning of the string
		if i == 0 && (char == '+' || char == '-') {
			if char == '-' {
				sign = -1
			}
			continue
		}

		// Check if the character is a digit
		if char >= '0' && char <= '9' {
			digit := int(char - '0')
			result = result*10 + digit
		} else {
			// If a non-digit character is encountered, return 0
			fmt.Println("There is one or more non-digit character(s) within the dataset. This program will interpret them as zero(es)")
			return 0
		}
	}

	return result * sign
}