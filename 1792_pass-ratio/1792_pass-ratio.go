package leetcode

import (
	"github.com/emirpasic/gods/trees/binaryheap"
)

type class struct {
	ratio, rInc float64
	pass, total int
}

func classComparator(a, b interface{}) int {
	ak := a.(*class)
	bk := b.(*class)
	switch {
	case ak.rInc < bk.rInc:
		return 1
	case ak.rInc > bk.ratio:
		return -1
	default:
		return 0
	}
}

func initClass(p, t int) *class {
	c := &class{}
	c.pass = p
	c.total = t
	c.ratio = float64(c.pass) / float64(c.total)
	c.rInc = float64(c.total-c.pass) / float64(c.total*(c.total+1))
	return c
}

func (c *class) addExtra() {
	c.pass++
	c.total++
	c.ratio = float64(c.pass) / float64(c.total)
	c.rInc = float64(c.total-c.pass) / float64(c.total*(c.total+1))
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	heap := binaryheap.NewWith(classComparator)
	for _, c := range classes {
		heap.Push(initClass(c[0], c[1]))
	}
	for i := 0; i < extraStudents; i++ {
		c, _ := heap.Pop()
		c.(*class).addExtra()
		heap.Push(c)
	}
	var sum float64
	for _, c := range heap.Values() {
		sum += c.(*class).ratio
	}
	return sum / float64(heap.Size())
}
