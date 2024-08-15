package calculations

import "testing"

func TestLinearRegression(t *testing.T) {
	type args struct {
		y []float64
	}
	tests := []struct {
		name  string
		args  args
		wantM float64
		wantC float64
	}{
		{
			name:  "Simple positive slope",
			args:  args{y: []float64{1, 2, 3, 4, 5}},
			wantM: 1,
			wantC: 1,
		},
		{
			name:  "Simple negative slope",
			args:  args{y: []float64{5, 4, 3, 2, 1}},
			wantM: -1,
			wantC: 5,
		},
		{
			name:  "Horizontal line",
			args:  args{y: []float64{3, 3, 3, 3, 3}},
			wantM: 0,
			wantC: 3,
		},
		{
			name:  "Vertical shift",
			args:  args{y: []float64{5, 5, 5, 5, 5}},
			wantM: 0,
			wantC: 5,
		},
		{
			name:  "Random points",
			args:  args{y: []float64{2, 4, 5, 4, 5}},
			wantM: 0.6,
			wantC: 2.8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, gotC := LinearRegression(tt.args.y)
			if gotM != tt.wantM {
				t.Errorf("LinearRegression() gotM = %v, want %v", gotM, tt.wantM)
			}
			if gotC != tt.wantC {
				t.Errorf("LinearRegression() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
