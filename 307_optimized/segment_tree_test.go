package leetcode

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	st := Constructor([]int{1, 3, 5, 3, 14, 0, 11, 7, 1})
	if st.SumRange(0, 2) != 9 {
		t.Fail()
	}
	if st.SumRange(2, 3) != 8 {
		t.Fail()
	}
	if st.SumRange(2, 6) != 33 {
		t.Fail()
	}
	st.Update(5, 2)
	if st.SumRange(2, 6) != 35 {
		t.Fail()
	}
}
