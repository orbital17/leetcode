package leetcode

import (
	"fmt"
	"testing"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{"basic", []int{2, 0, 2, 1, 1, 0}, "[0 0 1 1 2 2]"},
		{"basic", []int{2, 1, 2, 1, 1, 0}, "[0 1 1 1 2 2]"},
		{"basic", []int{0}, "[0]"},
		{"basic", []int{1}, "[1]"},
		{"basic", []int{2}, "[2]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.nums)
			if fmt.Sprint(tt.nums) != tt.want {
				t.Log(tt.nums)
				t.Fail()
			}
		})
	}
}
