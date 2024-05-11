package main

import (
	"fmt"
	"os"
	"math-skills/functions"
)

func main(){

	if len(os.Args) != 2{
		fmt.Println("Please enter exactly 2 arguments!")
	}

	input := os.Args[1]
	inputFile, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Average:",functions.Average(inputFile))
	fmt.Println("Median:",functions.Median(inputFile))
	fmt.Println("Variance:",functions.Variance(inputFile))
	fmt.Println("Standard Deviation:",functions.StandardDev(inputFile))

}
