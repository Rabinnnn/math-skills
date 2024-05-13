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
/*** Uncomment if you want an error message displayed incase the dataset contains some invalid data. Remember to import "strings"
	inputFileLine := strings.Split(string(inputFile), "\n")
	inputStr := functions.RemoveInvalid(inputFileLine)

	if len(inputFileLine) != len(inputStr){
		 fmt.Println("The dataset contains some invalid data (e.g letters, empty lines etc) that have been ignored!")
	} ***/

	fmt.Println("Average:",functions.Average(inputFile))
	fmt.Println("Median:",functions.Median(inputFile))
	fmt.Println("Variance:",functions.Variance(inputFile))
	fmt.Println("Standard Deviation:",functions.StandardDev(inputFile))
}
