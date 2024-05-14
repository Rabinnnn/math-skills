package main 

import(
	"testing"
	"fmt"
	"os"
	"math-skills/functions"
)

func TestAverage(t *testing.T){
	input, err := os.ReadFile("testData.txt")
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)	
	}

	expectedOutPut := int64(297)
	result := functions.Average(input)
	if result != expectedOutPut{
		t.Errorf("Expected %d but got %d", expectedOutPut, result)
	}
}

func TestMedian(t *testing.T){
	input, err := os.ReadFile("testData.txt")
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)	
	}

	expectedOutPut := int64(132)
	result := functions.Median(input)
	if result != expectedOutPut{
		t.Errorf("Expected %d but got %d", expectedOutPut, result)
	}
}

func TestVariance(t *testing.T){
	input, err := os.ReadFile("testData.txt")
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)	
	}

	expectedOutPut := int64( 130236)
	result := functions.Variance(input)
	if result != expectedOutPut{
		t.Errorf("Expected %d but got %d", expectedOutPut, result)
	}
}

func TestStandardDev(t *testing.T){
	input, err := os.ReadFile("testData.txt")
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)	
	}

	expectedOutPut := int64( 361)
	result := functions.StandardDev(input)
	if result != expectedOutPut{
		t.Errorf("Expected %d but got %d", expectedOutPut, result)
	}
}