package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	result := [][]int{}
	last := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= last[1] {
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			result = append(result, last)
			last = intervals[i]
		}
	}
	result = append(result, last)
	return result
}
