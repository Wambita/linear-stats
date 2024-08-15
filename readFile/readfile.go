package readfile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Readfile(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numslice := []float64{}
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("error converting string to float: %s\n", err)
			continue
		}
		numslice = append(numslice, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading from scanner: %s\n", err)
	}
	return numslice, nil
}

