package leetcode

import "container/heap"
import . "./lib"

type heapItem struct {
	Value int
	Index int
}

type kListsHeap []heapItem

func (h kListsHeap) Len() int           { return len(h) }
func (h kListsHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }
func (h kListsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *kListsHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(heapItem))
}

func (h *kListsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &kListsHeap{}
	for i, v := range lists {
		if v != nil {
			heap.Push(h, heapItem{v.Val, i})
			lists[i] = lists[i].Next
		}
	}
	dummy := &ListNode{}
	current := dummy
	for h.Len() != 0 {
		el := heap.Pop(h).(heapItem)
		current.Next = &ListNode{Val: el.Value}
		current = current.Next
		if lists[el.Index] != nil {
			heap.Push(h, heapItem{lists[el.Index].Val, el.Index})
			lists[el.Index] = lists[el.Index].Next
		}
	}
	return dummy.Next
}
