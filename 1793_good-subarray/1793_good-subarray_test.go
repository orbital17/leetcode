package leetcode

import "testing"

func Test_maximumScore(t *testing.T) {
	r := maximumScore([]int{1, 4, 3, 7, 4, 5}, 3)
	if r != 15 {
		t.Errorf("fail %d", r)
	}
}
