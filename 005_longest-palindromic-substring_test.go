package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	input := "babad"
	res := "bab"
	if longestPalindrome(input) != res {
		t.Fail()
	}
}

func Test_longestPalindrome2(t *testing.T) {

	input := "cbbd"
	res := "bb"
	if longestPalindrome(input) != res {
		t.Fail()
	}
}
func Test_longestPalindrome0(t *testing.T) {

	input := ""
	res := ""
	if longestPalindrome(input) != res {
		t.Fail()
	}
}
func Test_longestPalindrome3(t *testing.T) {

	input := "aaaaaaa"
	res := "aaaaaaa"
	if longestPalindrome(input) != res {
		t.Fail()
	}
}
