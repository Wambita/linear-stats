package calculations

import "testing"

func TestPearsonsCorrelationCoefficient(t *testing.T) {
	type args struct {
		y []float64
	}
	tests := []struct {
		name     string
		y        []float64
		expected float64
	}{
		{
			name:     "Perfect Positive Correlation",
			y:        []float64{1, 2, 3, 4, 5},
			expected: 0.9999999999999998, // when rounded its 1
		},
		{
			name:     "Perfect Negative Correlation",
			y:        []float64{5, 4, 3, 2, 1},
			expected: -0.9999999999999998,
		},

		{
			name:     "Single Value",
			y:        []float64{1},
			expected: 0.0, // With only one value, correlation is not meaningful
		},
		{
			name:     "All Same Values",
			y:        []float64{1, 1, 1, 1, 1},
			expected: 0.0, // No variance, hence no correlation
		},
		{
			name:     "Zero Variance in x",
			y:        []float64{1, 2, 3, 4, 5}, // x is just indices, correlation should be 1.0
			expected: 0.9999999999999998,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PearsonsCorrelationCoefficient(tt.y); got != tt.expected {
				t.Errorf("PearsonsCorrelationCoefficient() = %v, want %v", got, tt.expected)
			}
		})
	}
}
