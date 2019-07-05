package leetcode

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{"t", []int{1, 2, 3}, 4},
		{"t", []int{1}, 1},
		{"t", []int{2, 1, 5, 6, 2, 3}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
