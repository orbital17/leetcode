package leetcode

import "github.com/emirpasic/gods/sets/treeset"

func secondHighest(s string) int {
	set := treeset.NewWithIntComparator()
	for _, i := range s {
		if i >= '0' && i <= '9' {
			set.Add(int(i - '0'))
		}
	}
	if set.Size() < 2 {
		return -1
	}
	it := set.Iterator()
	it.Last()
	it.Prev()
	return it.Value().(int)
}

func secondHighest_old(s string) int {
	ai, bi := -1, -1
	var a, b byte
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if ai < 0 || s[i] > a {
				bi = ai
				b = a
				ai = i
				a = s[i]
				continue
			}
			if s[i] < a && (bi < 0 || s[i] > b) {
				bi = i
				b = s[i]
			}
		}
	}
	if bi < 0 {
		return bi
	}
	return int(b - '0')
}
