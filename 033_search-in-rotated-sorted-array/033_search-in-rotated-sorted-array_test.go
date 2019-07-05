package leetcode

import "testing"

func Test_search033(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{[]int{}, 4}, -1},
		{"basic", args{[]int{4}, 4}, 0},
		{"basic", args{[]int{4, 5, 6}, 4}, 0},
		{"basic", args{[]int{4, 5, 6}, 6}, 2},
		{"basic", args{[]int{4, 5, 6, 0, 1, 2, 3}, 2}, 5},
		{"basic", args{[]int{4, 5, 6, 0, 1, 2, 3}, 0}, 3},
		{"basic", args{[]int{4, 5, 6, 0, 1, 2, 3}, 6}, 2},
		{"basic", args{[]int{4, 5, 6, 0, 1, 2, 3}, 4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search033(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
