package leetcode

import "testing"

func Test_getNumberOfBacklogOrders(t *testing.T) {
	orders := [][]int{
		{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1},
	}
	r := getNumberOfBacklogOrders(orders)
	if r != 999999984 {
		t.Errorf("fail")
	}
}
