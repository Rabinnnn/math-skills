package functions

import(
	"math"
	"fmt"
)

func StandardDev(inputFile []byte) int64{
	n := Variance(inputFile)

	if n == 0 || n == 1 {
        return n
    }

	result := (n/2)
	fmt.Println("result:", result)

	for i := 0; i < int(result); i++{
		//Get the square root 
			result = result - (result*result-n)/(2*result)	
	}
	output := math.Round(float64(result))

return int64(output)
}
