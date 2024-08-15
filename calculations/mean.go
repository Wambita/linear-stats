package calculations

func mean(numSlice []float64) float64 {
	sum := 0.0
	for _, num := range numSlice {
		sum += num
	}
	return sum / float64(len(numSlice))
}
