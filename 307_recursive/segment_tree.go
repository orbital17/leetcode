package leetcode

type NumArray struct {
	data []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := 1
	for n < len(nums) {
		n <<= 1
	}
	res := NumArray{
		n:    n,
		data: make([]int, n<<1),
	}
	copy(res.data[n:], nums)
	for i := n - 1; i > 0; i-- {
		res.data[i] = res.data[i<<1] + res.data[i<<1|1]
	}
	return res
}

func (this *NumArray) Update(i int, val int) {
	n := this.n
	c := n + i
	this.data[c] = val
	for c > 0 {
		c >>= 1
		this.data[c] = this.data[c<<1] + this.data[c<<1|1]
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sumRangeRec(0, this.n-1, i, j, 1)
}

func (this *NumArray) sumRangeRec(l, r, ql, qr, index int) (res int) {
	if l == ql && r == qr {
		return this.data[index]
	} else {
		m := (l + r + 1) / 2
		if ql < m {
			res += this.sumRangeRec(l, m-1, ql, min(qr, m-1), index<<1)
		}
		if qr >= m {
			res += this.sumRangeRec(m, r, max(ql, m), qr, index<<1|1)
		}
		return res
	}
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
