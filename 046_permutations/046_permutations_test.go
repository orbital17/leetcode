package leetcode

import (
	"reflect"
	"testing"
)

func Test_nPerm(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"2", args{2}, [][]int{{0, 1}, {1, 0}}},
		{"3", args{3}, [][]int{
			{0, 1, 2},
			{0, 2, 1},
			{1, 0, 2},
			{1, 2, 0},
			{2, 0, 1},
			{2, 1, 0},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nPerm(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nPerm() = %v, want %v", got, tt.want)
			}
		})
	}
}
