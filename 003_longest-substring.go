package leetcode

//O(n*k) - number of symbols
func lengthOfLongestSubstring(s string) int {
	var i, j, r, max int
	for j != len(s) {
		j, r = findFrom(s, i)
		if j-i > max {
			max = j - i
		}
		i = r + 1
	}
	return max
}

func findFrom(s string, i int) (jr, r int) {
	m := make([]int, 128)
	for j := i; j < len(s); j++ {
		mi := int(s[j])
		if m[mi] == 0 {
			m[mi] = j + 1
		} else {
			return j, m[mi] - 1
		}
	}
	return len(s), 0
}
