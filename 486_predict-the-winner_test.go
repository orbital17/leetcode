package leetcode

import "testing"

func TestPredictTheWinner(t *testing.T) {
	nums := []int{1, 5, 2}
	res := false
	if res != PredictTheWinner(nums) {
		t.Fail()
	}
}

func TestPredictTheWinner2(t *testing.T) {
	nums := []int{1, 5, 233, 7}
	res := true
	if res != PredictTheWinner(nums) {
		t.Fail()
	}
}
