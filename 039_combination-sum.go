package leetcode

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	result := &[][]int{}
	sort.IntSlice(candidates).Sort()
	calcCombinationSum(result, []int{}, target, candidates)
	return *result
}

func calcCombinationSum(result *[][]int, current []int, target int, candidates []int) {
	for i, v := range candidates {
		switch {
		case v < target:
			calcCombinationSum(result, append(current, v), target-v, candidates[i:])
		case v == target:
			r := make([]int, len(current)+1)
			copy(r, current)
			r[len(current)] = v
			*result = append(*result, r)
		default:
			break
		}
	}
}
