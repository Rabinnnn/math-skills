package main 

import(
	"testing"
	"fmt"
	"os"
	"math-skills/functions"
	"bufio"
	"strconv"
	"log"
	"math"
)


func TestAverage(t *testing.T){
	inputFile, err := os.Open("testData.txt")
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
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatalf("Error %s \n", err)
		}
		numbers = append(numbers, num)

	}
	expectedOutPut := float64(149)
	result := math.Round(functions.Average(numbers))
	if result != expectedOutPut{
		t.Errorf("Expected %.0f but got %.0f", expectedOutPut, result)
	}
}

func TestMedian(t *testing.T){
	inputFile, err := os.Open("testData.txt")
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
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatalf("Error %s \n", err)
		}
		numbers = append(numbers, num)

	}
	expectedOutPut := float64(149)
	result := math.Round(functions.Median(numbers))
	if result != expectedOutPut{
		t.Errorf("Expected %.0f but got %.0f", expectedOutPut, result)
	}
}

func TestVariance(t *testing.T){
	inputFile, err := os.Open("testData.txt")
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
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatalf("Error %s \n", err)
		}
		numbers = append(numbers, num)

	}

	expectedOutPut := float64(791)
	result := math.Round(functions.Variance(numbers))
	if result != expectedOutPut{
		t.Errorf("Expected %.0f but got %.0f", expectedOutPut, result)
	}
}

func TestStandardDev(t *testing.T){
	inputFile, err := os.Open("testData.txt")
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
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatalf("Error %s \n", err)
		}
		numbers = append(numbers, num)

	}

	expectedOutPut := float64(28)
	result := math.Round(functions.StandardDev(numbers))
	if result != expectedOutPut{
		t.Errorf("Expected %.0f but got %.0f", expectedOutPut, result)
	}
}