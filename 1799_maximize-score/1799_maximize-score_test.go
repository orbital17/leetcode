package leetcode

import "testing"

func Test_maxScore(t *testing.T) {
	r := maxScore([]int{2, 3, 4, 6})
	if r != 8 {
		t.Errorf("fail")
	}
	r = maxScore([]int{925612, 737192, 813525, 707250})
	if r != 154 {
		t.Errorf("fail")
	}
}
