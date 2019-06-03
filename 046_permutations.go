package leetcode

func nPerm(n int) [][]int {
	if n == 0 {
		return [][]int{{}}
	}
	prevPerm := nPerm(n - 1)
	result := make([][]int, 0, len(prevPerm)*n)
	for i := 0; i < n; i++ {
		for _, vs := range prevPerm {
			r := make([]int, n)
			r[0] = i
			for j, v := range vs {
				r[j+1] = v
				if v >= i {
					r[j+1]++
				}
			}
			result = append(result, r)
		}
	}
	return result
}

func permute(nums []int) [][]int {
	result := nPerm(len(nums))
	for _, vi := range result {
		for j, vj := range vi {
			vi[j] = nums[vj]
		}
	}
	return result
}
