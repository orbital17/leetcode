package leetcode

import (
	"testing"
	"reflect"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := []int{0, 1}
	if !reflect.DeepEqual(res, twoSum(nums, target)) {
		t.Fail()
	}
}
