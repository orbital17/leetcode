package leetcode

import "testing"

func Test_maxValue(t *testing.T) {
	r := maxValue(8, 0, 15)
	if r != 4 {
		t.Errorf("fail")
	}
}
