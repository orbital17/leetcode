package leetcode

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	p := 1
	var first, middle, last int
	var firstSet, pSet, lastSet bool
	resetP := func() {
		if pSet && p > max {
			max = p
		}
		p = 1
		pSet = false
	}
	check := func(n int) {
		if n > max {
			max = n
		}
	}
	end := func() {
		if !firstSet {
			resetP()
		} else {
			check(first * middle)
			if lastSet {
				check(middle * last * p)
				check(first * middle * last * p)
			}
			resetP()
			firstSet = false
			lastSet = false
		}
	}
	for _, v := range nums {
		switch {
		case v < 0 && !firstSet:
			first = p * v
			middle = 1
			last = 1
			firstSet = true
			resetP()
		case v < 0 && firstSet:
			middle *= p * last
			last = v
			lastSet = true
			resetP()
		case v > 0:
			p *= v
			pSet = true
		case v == 0:
			check(0)
			end()
		}
	}
	end()
	return max
}
