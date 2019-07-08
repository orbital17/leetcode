package leetcode

import (
	"testing"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		a []int
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t", args{[]int{1, 2, 3}, 0}, 0},
		{"t", args{[]int{1, 2, 3}, 1}, 0},
		{"t", args{[]int{1, 2, 3}, 2}, 1},
		{"t", args{[]int{1, 2, 3}, 3}, 2},
		{"t", args{[]int{1, 2, 3}, 4}, 3},
		{"t", args{[]int{1}, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
