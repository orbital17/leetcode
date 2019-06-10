package leetcode

import (
	"fmt"
	"testing"
)

func Test_subsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{"basic", []int{}, "[[]]"},
		{"basic", []int{1}, "[[1] []]"},
		{"basic", []int{1, 2}, "[[2] [] [1 2] [1]]"},
		{"basic", []int{1, 2, 3}, "[[3] [] [2 3] [2] [1 3] [1] [1 2 3] [1 2]]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.nums); fmt.Sprint(got) != tt.want {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
