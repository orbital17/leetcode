package leetcode

import "testing"

func Test_getMoneyAmount(t *testing.T) {
	n, result := 11, 18
	if getMoneyAmount(n) != result {
		t.Fail()
	}
}
