package leetcode

import "testing"

func Test_coinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{"t", []int{1, 2, 5}, 11, 3},
		{"t", []int{3, 4, 16}, 17, 5},
		{"t", []int{3}, 17, -1},
		{"t", []int{3}, 3, 1},
		{"t", []int{}, 0, 0},
		{"t", []int{}, 1, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.coins, tt.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
