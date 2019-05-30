package leetcode

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	points := [][]int{{1, 3}, {-2, 2}}
	K := 1
	res := [][]int{{-2, 2}}
	actual := kClosest(points, K)
	if !reflect.DeepEqual(res, actual) {
		t.Fail()
	}
}

func Test_kClosest2(t *testing.T) {
	points := [][]int{{3, 3}, {5, -1}, {-2, -4}}
	K := 2
	res := [][]int{{-2, -4}, {3, 3}}
	actual := kClosest(points, K)
	if !reflect.DeepEqual(res, actual) {
		t.Fail()
	}
}
