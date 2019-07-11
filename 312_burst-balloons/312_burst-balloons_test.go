package leetcode

import "testing"

func Test_maxCoins(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t", []int{5, 3}, 20},
		{"t", []int{3}, 3},
		{"t", []int{3, 1, 5, 8}, 167},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.nums); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
