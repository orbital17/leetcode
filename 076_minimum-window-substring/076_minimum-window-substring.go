package leetcode

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	count := make(map[byte]int)
	thash := make(map[byte]int)
	for _, v := range t {
		count[byte(v)] = 0
		if _, ok := thash[byte(v)]; ok {
			thash[byte(v)]++
		} else {
			thash[byte(v)] = 1
		}
	}
	l, r := 0, 0
	unique := 0
	minL, minR := 0, len(s)+1
	for {
		if unique < len(t) {
			if r == len(s) {
				break
			}
			r++
			v := s[r-1]
			c, ok := count[v]
			if !ok {
				continue
			}
			count[v]++
			if c < thash[v] {
				unique++
				if unique == len(t) && r-l < minR-minL {
					minL, minR = l, r
				}
			}
		} else {
			l++
			v := s[l-1]
			c, ok := count[v]
			if !ok {
				continue
			}
			count[v]--
			if c == thash[v] {
				unique--
				if r-l+1 < minR-minL {
					minL, minR = l-1, r
				}
			}
		}
	}
	if minR == len(s)+1 {
		return ""
	}
	return s[minL:minR]
}
