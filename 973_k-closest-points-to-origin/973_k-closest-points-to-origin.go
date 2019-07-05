package leetcode

import (
	"container/heap"
)

type kClosestHeap struct {
	points [][]int
	K      int
}

func (h kClosestHeap) Len() int { return h.K }
func (h kClosestHeap) Less(i, j int) bool {
	return squareSum(h.points[i]) > squareSum(h.points[j])
}
func (h kClosestHeap) Swap(i, j int) { h.points[i], h.points[j] = h.points[j], h.points[i] }

func (h *kClosestHeap) Push(x interface{}) {
	panic("not implemented")
}
func (h *kClosestHeap) Pop() interface{} {
	panic("not implemented")
}

func squareSum(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

func kClosest(points [][]int, K int) [][]int {
	h := &kClosestHeap{
		points, K,
	}
	heap.Init(h)
	for i := K; i < len(points); i++ {
		if squareSum(points[i]) < squareSum(points[0]) {
			h.Swap(0, i)
			heap.Fix(h, 0)
		}
	}
	return points[0:K]
}
