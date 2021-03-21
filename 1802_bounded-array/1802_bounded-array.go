package leetcode

func sum(n int, index int, k int) int {
	m := 0
	sum := 0
	if k > index+1 {
		m = k - index
	}
	sum += (m + k) * (k - m + 1) / 2

	index = n - 1 - index
	m = 0
	if k > index+1 {
		m = k - index
	}
	sum += (m + k) * (k - m + 1) / 2
	sum -= k
	return sum
}

func maxValue(n int, index int, maxSum int) int {
	maxSum -= n

	l, r := 0, 1000000001
	for r > l+1 {
		k := (r + l) / 2
		s := sum(n, index, k)
		if s > maxSum {
			r = k
			continue
		}
		if s < maxSum {
			l = k
			continue
		}
		if s == maxSum {
			return k + 1
		}
	}
	if sum(n, index, r) <= maxSum {
		return r + 1
	}
	return l + 1
}
