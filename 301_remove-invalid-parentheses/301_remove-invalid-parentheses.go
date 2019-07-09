package leetcode

func removeInvalidParentheses(s string) []string {
	n := len(s)
	oldQ := [][]int{}
	newQ := [][]int{}
	oldQ = append(oldQ, []int{-1})
	results := [][]int{}
	for {
		for _, v := range oldQ {
			valid, before := isValid(s, v)
			switch {
			case valid == 0:
				results = append(results, v)
			case valid > 0:
				last := v[len(v)-1]
				for i := last + 1; i < n; i++ {
					if s[i] == '(' {
						new := make([]int, len(v)+1)
						copy(new, v)
						new[len(v)] = i
						newQ = append(newQ, new)
					}
				}
			case valid < 0:
				last := v[len(v)-1]
				for i := last + 1; i < before; i++ {
					if s[i] == ')' {
						new := make([]int, len(v)+1)
						copy(new, v)
						new[len(v)] = i
						newQ = append(newQ, new)
					}
				}
			}
		}
		oldQ = newQ
		newQ = [][]int{}
		if len(results) != 0 {
			break
		}
	}
	resHash := make(map[string]bool)
	for _, v := range results {
		cur := []rune{}
		j := 1
		for i, c := range s {
			if j >= len(v) || i != v[j] {
				cur = append(cur, c)
			} else {
				j++
			}
		}
		resHash[string(cur)] = true
	}
	r := []string{}
	for k := range resHash {
		r = append(r, k)
	}
	return r
}

func isValid(s string, mask []int) (res, before int) {
	c := 0
	j := 1
	for i, v := range s {
		switch {
		case j < len(mask) && i == mask[j]:
			j++
			continue
		case v == '(':
			c++
		case v == ')':
			c--
			if c < 0 {
				return c, i + 1
			}
		}
	}
	return c, len(s)
}
