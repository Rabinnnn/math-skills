package functions

func StandardDev(inputFile []byte) int{
	n := Variance(inputFile)
	result := 0

	for i := 0; i < n; i++{
		if (i*i) == n{ //Get the square root 
			result = i
		}
	}
return result
}