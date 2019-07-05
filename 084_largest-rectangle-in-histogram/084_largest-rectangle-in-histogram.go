package leetcode

type List84 struct {
	next *List84
	n    int
	d    int
}

func largestRectangleArea(heights []int) int {
	var ds, cur *List84
	var max, d int
	for _, v := range heights {
		d = 0
		for cur = ds; cur != nil && cur.n > v; cur = cur.next {
			d += cur.d
			if cur.n*d > max {
				max = cur.n * d
			}
		}
		if cur == nil || v > cur.n {
			ds = &List84{cur, v, d + 1}
		} else { //cur.n == v
			cur.d += d + 1
			ds = cur
		}
	}
	d = 0
	for cur = ds; cur != nil; cur = cur.next {
		d += cur.d
		if cur.n*d > max {
			max = cur.n * d
		}
	}
	return max
}
