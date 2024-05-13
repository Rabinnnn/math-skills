package functions

func RemoveInvalid(input []string) []string{
	var output []string

	for _, num := range input{
		if StrToInt(num) >= 0 && num != "" {
			output = append(output, num)
		} 
		
	}
	return output
}