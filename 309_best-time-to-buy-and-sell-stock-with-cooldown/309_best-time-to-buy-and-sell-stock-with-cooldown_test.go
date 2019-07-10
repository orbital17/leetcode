package cooldown

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"t", []int{1, 2, 3, 0, 2}, 3},
		{"t", []int{}, 0},
		{"t", []int{1}, 0},
		{"t", []int{1, 2}, 1},
		{"t", []int{2, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
