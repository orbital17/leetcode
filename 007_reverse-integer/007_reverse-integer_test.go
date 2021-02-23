package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	num := 123
	res := 321
	if res != reverse(num) {
		t.Fail()
	}
}

func Test_reverseNegatives(t *testing.T) {
	num := -123
	res := -321
	if res != reverse(num) {
		t.Fail()
	}
}

func Test_reverseOverflow(t *testing.T) {
	num := 2147483638
	res := 0
	actual := reverse(num)
	if res != actual {
		t.Fail()
	}
}
