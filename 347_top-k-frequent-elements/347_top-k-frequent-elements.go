package leetcode

import "container/heap"

type pHeap struct {
	fsmap map[int]int
	p     []int
	K     int
}

func (h pHeap) Len() int { return h.K }
func (h pHeap) Less(i, j int) bool {
	return h.fsmap[h.p[i]] < h.fsmap[h.p[j]]
}
func (h pHeap) Swap(i, j int) { h.p[i], h.p[j] = h.p[j], h.p[i] }

func (h *pHeap) Push(x interface{}) {
	panic("not implemented")
}
func (h *pHeap) Pop() interface{} {
	panic("not implemented")
}

func topKFrequent(nums []int, k int) []int {
	fsmap := make(map[int]int)
	for _, v := range nums {
		fsmap[v]++
	}
	fs := make([]int, k)
	i := 0
	h := &pHeap{fsmap, fs, k}
	for el, f := range fsmap {
		if i < k {
			fs[i] = el
		}
		if i == k {
			heap.Init(h)
		}
		if i >= k && f > fsmap[fs[0]] {
			fs[0] = el
			heap.Fix(h, 0)
		}
		i++
	}
	for i = 0; i < k/2; i++ {
		fs[i], fs[k-1-i] = fs[k-1-i], fs[i]
	}
	return fs
}
