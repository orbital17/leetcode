package leetcode

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.IntSlice(coins).Sort()
	var current int
	for i := 0; i < len(coins); i++ {
		if current+1 >= coins[i] {
			current += coins[i]
		} else {
			return current + 1
		}
	}
	return current + 1
}

func getMaximumConsecutive_fail(coins []int) int {
	sort.IntSlice(coins).Sort()
	l := len(coins)
	maxn := 4*10000 + 3
	var d [][]bool
	d = [][]bool{make([]bool, l+1)}
	for i := 0; i <= l; i++ {
		d[0][i] = true
	}
	for n := 1; n < maxn*maxn; n++ {
		d = append(d, make([]bool, l+1))
		for i := 1; i <= l; i++ {
			if d[n][i-1] || n-coins[i-1] < 0 {
				d[n][i] = d[n][i-1]
			} else {
				d[n][i] = d[n-coins[i-1]][i-1]
			}
		}
		if !d[n][l] {
			return n
		}
	}

	return maxn
}
