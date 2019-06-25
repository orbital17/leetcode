package leetcode

import (
	"testing"
)

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t", []int{1, 2, 3, 1}, 4},
		{"t", []int{2, 7, 9, 3, 1}, 12},
		{"t", []int{2, 7, 1, 1, 9, 3, 1}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
