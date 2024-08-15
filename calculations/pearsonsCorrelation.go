package calculations

import "math"

// formula sum((xi - meanX)(yi-meany)) / squareroot((sum(xi-meanx)(sum(yi-meany))))
func PearsonsCorrelationCoefficient(y []float64) float64 {
	x := make([]float64, len(y))

	for i := range y {
		x[i] = float64(i)
	}

	// means
	meanX := Mean(x)
	meanY := Mean(y)

	var sumXY, squareX, squareY float64
	for i, yi := range y {
		xi := float64(i)
		sumXY += (xi - meanX) * (yi - meanY)
		squareX += (xi - meanX) * (xi - meanX)
		squareY += (yi - meanY) * (yi - meanY)

	}

	// coefficient
	rootXy := math.Sqrt(squareX * squareY)

	if rootXy == 0 {
		return 0
	}

	correlation := sumXY / rootXy
	return correlation
}
