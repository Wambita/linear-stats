package main

import (
	"fmt"
	"os"

	"linear-stats/calculations"
	readfile "linear-stats/readFile"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run . data.txt")
		os.Exit(1)
	}
	filename := os.Args[1]
	numslice, err := readfile.Readfile(filename)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	m, c := calculations.LinearRegression(numslice)
	correlation := calculations.PearsonsCorrelationCoefficient(numslice)

	fmt.Printf("Linear Regression Line: y = %fx + %f\n", m, c)
	fmt.Printf("Pearson Correlation Coefficient: %f\n", correlation)
	// Linear Regression Line: y = <value>x + <value>
	// Pearson Correlation Coefficient: <value>
}
