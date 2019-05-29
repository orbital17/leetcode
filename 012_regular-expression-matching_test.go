package leetcode

import "testing"

func Test_isMatch(t *testing.T) {
	text := "addde"
	pattern := "ad*e"
	expected := true
	actual := isMatch(text, pattern)
	if expected != actual {
		t.Fail()
	}
}

func Test_isMatch2(t *testing.T) {
	text := ""
	pattern := "d*"
	expected := true
	actual := isMatch(text, pattern)
	if expected != actual {
		t.Fail()
	}
}

func Test_isMatch3(t *testing.T) {
	text := "ab"
	pattern := ".*c"
	expected := false
	actual := isMatch(text, pattern)
	if expected != actual {
		t.Fail()
	}
}

func Test_isMatch4(t *testing.T) {
	text := "ab"
	pattern := ".*"
	expected := true
	actual := isMatch(text, pattern)
	if expected != actual {
		t.Fail()
	}
}

func Test_isMatch5(t *testing.T) {
	text := ""
	pattern := ""
	expected := true
	actual := isMatch(text, pattern)
	if expected != actual {
		t.Fail()
	}
}
