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

	// format values to required decimal places
	mformatted := fmt.Sprintf("%.*f", 6, m)
	cformatted := fmt.Sprintf("%.*f", 6, c)
	correlationFormatted := fmt.Sprintf("%.*f", 10, correlation)

	fmt.Printf("Linear Regression Line: y = %sx + %s\n", mformatted, cformatted)
	fmt.Printf("Pearson Correlation Coefficient: %s\n", correlationFormatted)
}
