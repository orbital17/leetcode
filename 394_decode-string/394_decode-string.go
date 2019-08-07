package leetcode

type Pair struct {
	n int
	s []byte
}

func decodeString(s string) string {
	b := []byte(s)
	stack := []Pair{
		{1, []byte{}},
	}
	i := 0
	for i < len(b) {
		v := b[i]
		if v >= '0' && v <= '9' {
			n := 0
			for b[i] != '[' {
				n *= 10
				n += int(b[i] - '0')
				i++
			}
			stack = append(stack, Pair{n, []byte{}})
			i++
			continue
		}
		last := &stack[len(stack)-1]
		if v != ']' {
			last.s = append(last.s, v)
		} else {
			prev := &stack[len(stack)-2]
			for j := 0; j < last.n; j++ {
				for _, lv := range last.s {
					prev.s = append(prev.s, lv)
				}
			}
			stack = stack[0 : len(stack)-1]
		}
		i++
	}
	return string(stack[0].s)
}
