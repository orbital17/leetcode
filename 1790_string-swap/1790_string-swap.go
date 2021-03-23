package leetcode

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	i, j := -1, -1
	for k := 0; k < len(s1); k++ {
		if s1[k] != s2[k] {
			if i < 0 {
				i = k
				continue
			}
			if j < 0 {
				j = k
				continue
			}
			return false
		}
	}
	if j < 0 {
		return i < 0
	}
	return s1[i] == s2[j] && s1[j] == s2[i]
}
