package leetcode

import "testing"

func Test_maxAverageRatio(t *testing.T) {
	r := maxAverageRatio([][]int{
		{1, 2}, {3, 5}, {2, 2},
	}, 2)
	if r != 0.783333 {
		t.Errorf("%f", r)
	}
}
