package leetcode

import "testing"

func Test_convert(t *testing.T) {
	s := "PAYPALISHIRING"
	rows := 4
	expected := "PINALSIGYAHRPI"
	actual := convert(s, rows)
	if actual != expected {
		t.Fail()
	}
}

func Test_convert2(t *testing.T) {
	s := "PAYPALISHIRING"
	rows := 3
	expected := "PAHNAPLSIIGYIR"
	actual := convert(s, rows)
	if actual != expected {
		t.Fail()
	}
}
