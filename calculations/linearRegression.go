package calculations

func LinearRegression(y []float64) (m, c float64) {
	// y dependent variable
	// x indices, independent var
	// m = slope
	// intercept

	// generate x axis values
	x := make([]float64, len(y))
	for i := range y {
		x[i] = float64(i)
	}

	// get means
	meanX := Mean(x)
	meanY := Mean(y)

	// calculation of m and c
	// XYSum sum of products of deviations
	// X2Sum sum of squared deviations

	var XYSum, X2Sum float64

	for i, yi := range y {
		xi := float64(i)
		XYSum += (xi - meanX) * (yi - meanY)
		X2Sum += (xi - meanX) * (xi - meanX)
	}

	// slope (rate at whixh y chnages for each unit of x)
	//
	m = XYSum / X2Sum

	// where the regression line crosses the y axis
	// c = y - mx
	c = meanY - m*meanX
	return
}
