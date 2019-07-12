package leetcode

type NumArray struct {
	data []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
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
	j += 1 + this.n
	i += this.n
	res := 0
	for i < j {
		if i&1 > 0 {
			res += this.data[i]
			i++
		}
		if j&1 > 0 {
			j--
			res += this.data[j]
		}
		i >>= 1
		j >>= 1
	}
	return res
}
