package leetcode

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t", args{[]int{4, 1, 2, 3, 5, 4}}, 4},
		{"t", args{[]int{4, 1, 2, 3, 5, 5}}, 5},
		{"t", args{[]int{4, 1, 2, 3, 5, 6, 7, 4}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
