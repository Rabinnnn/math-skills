package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"math-skills/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter exactly 2 arguments!")
		return
	}

	input := os.Args[1]

	//Only accept .txt files
	if !strings.HasSuffix(input, ".txt") {
		fmt.Println("file format not supported: only .txt files are allowed")
		os.Exit(1)
	}

	inputFile, err := os.Open(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer inputFile.Close()
	
	//Read through each line, convert valid data to float64 then add them to a new slice
	scanner := bufio.NewScanner(inputFile)
	numbers := []float64{}
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if err != nil {
			log.Fatalf("Error %s \n", err)
		}
		numbers = append(numbers, num)



	}
	if len(numbers) == 0{
		fmt.Println("The file is empty!")
		return
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}

	fmt.Printf("Average: %.0f\n", math.Round(functions.Average(numbers)))
	fmt.Println("Median: ", math.Round(functions.Median(numbers)))
	fmt.Printf("Variance: %.0f\n", math.Round(functions.Variance(numbers)))
	fmt.Printf("Standard Deviation: %.0f\n", math.Round(functions.StandardDev(numbers)))
}
