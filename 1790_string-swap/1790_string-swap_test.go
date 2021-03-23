package leetcode

import "testing"

func Test_areAlmostEqual(t *testing.T) {
	r := areAlmostEqual("bank", "kanb")
	if r != true {
		t.Error()
	}
	r = areAlmostEqual("", "")
	if r != true {
		t.Error()
	}
	r = areAlmostEqual("ab", "ab")
	if r != true {
		t.Error()
	}
	r = areAlmostEqual("a", "b")
	if r != false {
		t.Error()
	}
	r = areAlmostEqual("ab", "bb")
	if r != false {
		t.Error()
	}
}
