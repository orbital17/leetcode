package leetcode

import . "../lib"

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var i, j int
	n, m := len(matrix), len(matrix[0])
	h := make([]int, m)
	left := make([]int, m)
	right := make([]int, m)
	Fill(right, m)
	var cl, cr, cs, maxs int
	for i = 0; i < n; i++ {
		cl = 0
		for j = 0; j < m; j++ {
			if matrix[i][j] == '0' {
				left[j] = 0
				cl = j + 1
				h[j] = 0
			} else {
				left[j] = Max(left[j], cl)
				h[j]++
			}
		}
		cr = m
		for j = m - 1; j >= 0; j-- {
			if matrix[i][j] == '0' {
				right[j] = m
				cr = j
			} else {
				right[j] = Min(right[j], cr)
				cs = (right[j] - left[j]) * h[j]
				maxs = Max(cs, maxs)
			}
		}
	}
	return maxs
}

type List85 struct {
	next *List85
	n    byte
	d    byte
}

func largestRectangleArea85(heights []byte) byte {
	var ds, cur *List85
	var max, d byte
	for _, v := range heights {
		d = 0
		for cur = ds; cur != nil && cur.n > v; cur = cur.next {
			d += cur.d
			if cur.n*d > max {
				max = cur.n * d
			}
		}
		if cur == nil || v > cur.n {
			ds = &List85{cur, v, d + 1}
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

func maximalRectangle_reuse84(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var c byte
	for j := range matrix[0] {
		c = 0
		for i := range matrix {
			if matrix[i][j] == '1' {
				c++
			} else {
				c = 0
			}
			matrix[i][j] = c
		}
	}
	max := 0
	for i := range matrix {
		m := int(largestRectangleArea85(matrix[i]))
		if m > max {
			max = m
		}
	}
	return max
}
