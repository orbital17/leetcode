package leetcode

type node struct {
	zero, one *node
	count     int
}

const max = 15

func (root *node) insert(n int) {
	cur := root
	for i := max; i >= 0; i-- {
		if cur.zero == nil {
			cur.zero = &node{}
		}
		if cur.one == nil {
			cur.one = &node{}
		}
		bit := (n >> i) & 1
		next := cur.zero
		if bit > 0 {
			next = cur.one
		}
		next.count++
		cur = next
	}
}

func (root *node) countPairsVal(val, high int) int {
	cur := root
	var sum int
	for i := max; i >= 0; i-- {
		hbit := (high >> i) & 1
		vbit := (val >> i) & 1
		xor := vbit | hbit - vbit&hbit
		next := cur.zero
		other := cur.one
		if xor > 0 {
			next, other = other, next
		}
		if hbit > 0 && other != nil {
			sum += other.count
		}
		if next == nil {
			break
		}
		cur = next
	}
	return sum
}

func countPairs(nums []int, low int, high int) int {
	root := &node{}
	var sum int
	for _, n := range nums {
		sum += root.countPairsVal(n, high+1) - root.countPairsVal(n, low)
		root.insert(n)
	}
	return sum
}
