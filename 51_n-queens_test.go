package leetcode

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]string
	}{
		{"t", 1, [][]string{{"Q"}}},
		{"t", 3, [][]string{}},
		{"t", 0, [][]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
