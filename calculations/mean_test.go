package calculations

import "testing"

func TestMean(t *testing.T) {
	type args struct {
		numslice []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Single element",
			args: args{numslice: []float64{10.0}},
			want: 10,
		},
		{
			name: "Single element",
			args: args{numslice: []float64{10.0}},
			want: 10.0,
		},
		{
			name: "Multiple elements",
			args: args{numslice: []float64{1.0, 2.0, 3.0, 4.0, 5.0}},
			want: 3.0,
		},
		{
			name: "Mixed positive and negative",
			args: args{numslice: []float64{-1.0, -2.0, 3.0, 4.0}},
			want: 1.0,
		},
		{
			name: "All negative numbers",
			args: args{numslice: []float64{-1.0, -2.0, -3.0, -4.0, -5.0}},
			want: -3.0,
		},
		{
			name: "All same numbers",
			args: args{numslice: []float64{2.0, 2.0, 2.0, 2.0, 2.0}},
			want: 2.0,
		},
		{
			name: "Floating point numbers",
			args: args{numslice: []float64{1.5, 2.5, 3.5, 4.5}},
			want: 3.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mean(tt.args.numslice); got != tt.want {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}
