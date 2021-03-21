package leetcode

import "testing"

func Test_maxAscendingSum(t *testing.T) {
	r := maxAscendingSum([]int{1})
	if r != 1 {
		t.Errorf("fail")
	}
	r = maxAscendingSum([]int{1, 2, 3})
	if r != 6 {
		t.Errorf("fail")
	}
	r = maxAscendingSum([]int{1, 2, 3, 1, 7})
	if r != 8 {
		t.Errorf("fail")
	}
	r = maxAscendingSum([]int{1, 2, 9, 1, 7})
	if r != 12 {
		t.Errorf("fail")
	}
}
