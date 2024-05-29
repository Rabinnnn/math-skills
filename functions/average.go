package functions

func Average(inputFile []float64) float64 {

	var count, sum float64

	for i := 0; i < len(inputFile); i++ {
		count += 1
		sum += inputFile[i]

	}

	average := sum / count

	return average
}
