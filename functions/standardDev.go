package functions

import(
	"math"
)

func StandardDev(inputFile []float64) float64{
	n := Variance(inputFile)

	if n == 0 || n == 1 {
        return n
    }

	output := math.Sqrt(n)

return output
}
