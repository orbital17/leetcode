package leetcode

import "testing"

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t", []int{1, 2, 3}, 6},
		{"t", []int{1, 2, 0}, 2},
		{"t", []int{1, -1, 0}, 1},
		{"t", []int{-1, -1, 0}, 1},
		{"t", []int{-1, 0, -2}, 0},
		{"t", []int{-2, -3, -4}, 12},
		{"t", []int{-2, -3, -4, -5}, 120},
		{"t", []int{-2, -3, -4, -5, -1}, 120},
		{"t", []int{-2, -3, 0, -4, -5, -1}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
