# Linear-Stats

A Go program to compute  Linear Regression Line and Pearson Correlation Coefficient.

## Overview

This program reads numerical data from a file and calculates:

 + **Linear Regression Line**: Determines the best-fit line through the data points.
    
 + **Pearson Correlation Coefficient**: Measures the strength and direction of the linear relationship between two variables.

### Features

 - Reads numerical data from a file.
 - Computes Linear Regression Line (slope and intercept).
 - Computes Pearson Correlation Coefficient.
 - Outputs results with specified decimal precision.

## Installation

1. Install Go: Ensure you have Go installed on your system. You can download and install Go from the official Go website.

2. Clone the Repository: Clone the repository containing the linear-stats program. If you have a GitHub repository, you can use:

 ```  bash

git clone https://learn.zone01kisumu.ke/git/shfana/linear-stats.git
```


## Usage

1. Navigate to the Project `linear-stats`. folder:

```bash

cd linear-stats
```

2. Create a Data File: Save your data in a text file named data.txt in the root directory. For example:
```
189
113
121
114
145
110
```
3. Run the Program:


**Option1**

Run the program directly on the terminal by running the command `go run . ` followed by the data file `data.txt`

```sh

go run . data.txt
```
Output:

```bash

Linear Regression Line: y = 0.456789x + 123.456789
Pearson Correlation Coefficient: 0.9876543210
```

**option2**

1. Build the Program (optional): You can build the program to create an executable binary:

```sh

go build -o  linear-stats main.go
```


2.After building the executable, run it as:

```sh

    ./linear-stats data.txt
```
## Implementation

### Functions Used

#### Mean
 - **Purpose:** Computes the average value of a slice of float64 numbers.

 - **Signature:** func Mean(data []float64) float64

 - **Description:** Calculates the sum of all values in the slice and divides by the number of values to obtain the mean.

#### LinearRegression
- Purpose: Calculates the slope (m) and intercept (c) for the linear regression line.

- Signature: func LinearRegression(y []float64) (m, c float64)

- Description: Uses the least squares method to find the best-fit line for the given data points.

#### PearsonCorrelationCoefficient
- Purpose: Computes the Pearson Correlation Coefficient for the given data.
        
- Signature: func PearsonCorrelationCoefficient(y []float64) float64

- Description: Measures the linear correlation between two datasets, where one is the indices and the other is the given data.

#### ReadFile
- Purpose: Reads numeric data from a file and returns it as a slice of float64.
- Signature: func ReadFile(filename string) ([]float64, error)
- Description: Opens the file, reads each line, converts it to float64, and returns a slice of these values.

#### main
- Purpose: The entry point of the program.
- Description: Parses command-line arguments, reads data from the specified file, computes statistical metrics, and prints the results.

## Learning outcomes

This project help you learn about

- Statistical and Probabilities Calculations 

 [Linear Regression Line]('https://en.wikipedia.org/wiki/Linear_regression')

 [Pearson Correlation Coeffeicient]('https://en.wikipedia.org/wiki/Pearson_correlation_coefficient')


## Contributions

Contributions are welcome! If you have suggestions or improvements, please fork the repository and submit a pull request. Ensure that your contributions adhere to the project's coding standards and include appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Author

This project was developed by [Wambita Sheila Fana]('https://learn.zone01kisumu.ke/git/shfana').