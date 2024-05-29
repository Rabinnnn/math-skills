package functions

func Variance(inputFile []float64) float64 {
	mean := Average(inputFile)

	var sum, count float64

	//calculate variance
	for i := 0; i < len(inputFile); i++ {
		count++
		diff := inputFile[i] - mean
		diff2 := (diff * diff)
		sum += diff2
	}
	result := sum / count

	return result
}
