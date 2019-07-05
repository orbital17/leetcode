package leetcode

func longestPalindrome(s string) string {
	b := []byte(s)
	var r, max int
	var maxs []byte
	for i := range b {
		r = checkPalindrome(b, i, false)
		if r > max {
			max = r
			maxs = b[i-r/2 : i+r/2+1]
		}
		r = checkPalindrome(b, i, true)
		if r > max {
			max = r
			maxs = b[i+1-r/2 : i+r/2+1]
		}
	}
	return string(maxs)
}

func checkPalindrome(s []byte, i int, even bool) int {
	var l, r int
	if even {
		l, r = i, i+1
	} else {
		l, r = i-1, i+1
	}
	for {
		if l < 0 || r >= len(s) || s[l] != s[r] {
			return r - l - 1
		}
		l--
		r++
	}
}
