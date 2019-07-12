package leetcode

type NumArray struct {
	data []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := 1
	for n < len(nums) {
		n *= 2
	}
	res := NumArray{
		n:    n,
		data: make([]int, 2*n-1),
	}
	copy(res.data[n-1:], nums)
	for i := n - 2; i >= 0; i-- {
		res.data[i] = res.data[2*i+1] + res.data[2*i+2]
	}
	return res
}

func (this *NumArray) Update(i int, val int) {
	n := this.n
	c := n - 1 + i
	this.data[c] = val
	for c > 0 {
		c = (c - 1) / 2
		this.data[c] = this.data[2*c+1] + this.data[2*c+2]
	}
}

type t struct {
	l, r, ql, qr, index int
}

func (this *NumArray) SumRange(i int, j int) int {
	n := this.n
	que := []t{
		{0, n - 1, i, j, 0},
	}
	res := 0
	for len(que) != 0 {
		c := que[0]
		que = que[1:]
		if c.l == c.ql && c.r == c.qr {
			res += this.data[c.index]
		} else {
			m := (c.l + c.r + 1) / 2
			if c.ql < m {
				que = append(que, t{c.l, m - 1, c.ql, min(c.qr, m-1), 2*c.index + 1})
			}
			if c.qr >= m {
				que = append(que, t{m, c.r, max(c.ql, m), c.qr, 2*c.index + 2})
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
