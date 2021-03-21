package leetcode

import "testing"

func Test_countPairs(t *testing.T) {
	r := countPairs([]int{9, 8, 4, 2, 1}, 5, 14)
	if r != 8 {
		t.Error()
	}
}
