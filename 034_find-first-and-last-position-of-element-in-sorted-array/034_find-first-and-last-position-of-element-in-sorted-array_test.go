package leetcode

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"basic", args{[]int{}, 4}, []int{-1, -1}},
		{"basic", args{[]int{3}, 4}, []int{-1, -1}},
		{"basic", args{[]int{4}, 4}, []int{0, 0}},
		{"basic", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
		{"basic", args{[]int{1}, 0}, []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
