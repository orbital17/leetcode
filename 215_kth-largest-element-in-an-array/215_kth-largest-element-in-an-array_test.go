package leetcode

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t", args{[]int{1, 2, 3}, 2}, 2},
		{"t", args{[]int{3, 2, 1, 5, 6, 4}, 2}, 5},
		{"t", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4}, 4},
		{"t", args{[]int{1}, 1}, 1},
		{"t", args{[]int{-1, -1}, 2}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
