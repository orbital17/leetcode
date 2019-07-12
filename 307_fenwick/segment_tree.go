package leetcode

type NumArray struct {
	tree []int
	nums []int
}

func Constructor(nums []int) NumArray {
	r := NumArray{
		nums: nums,
		tree: make([]int, len(nums)+1),
	}
	for i, v := range nums {
		r.init(i, v)
	}
	return r
}

func (this *NumArray) init(i, val int) {
	for i += 1; i < len(this.tree); i += i & -i {
		this.tree[i] += val
	}
}

func (this *NumArray) Update(i int, val int) {
	this.init(i, val-this.nums[i])
	this.nums[i] = val
}

func (this *NumArray) prefixSum(i int) (r int) {
	for i += 1; i > 0; i -= i & -i {
		r += this.tree[i]
	}
	return
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.prefixSum(j) - this.prefixSum(i-1)
}
