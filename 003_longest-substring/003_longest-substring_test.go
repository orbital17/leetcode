package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	s := "AZazabcdefgabcdefghaa190!@#aa$%"
	res := 8
	actual := lengthOfLongestSubstring(s)
	if res != actual {
		t.Fail()
	}
}
