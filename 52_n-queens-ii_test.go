package leetcode

import "testing"

func Test_totalNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"t", 4, 2},
		{"t", 1, 1},
		{"t", 5, 10},
		{"t", 8, 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.n); got != tt.want {
				t.Errorf("totalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
