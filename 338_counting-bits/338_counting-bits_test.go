package leetcode

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want []int
	}{
		{"t", 9, []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2}},
		{"t", 1, []int{0, 1}},
		{"t", 0, []int{0}},
		{"t", 3, []int{0, 1, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
