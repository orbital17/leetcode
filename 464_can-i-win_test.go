package leetcode

import "testing"

func Test_canIWin(t *testing.T) {
	res := canIWin(20, 210)
	if res {
		t.Fail()
	}
}
