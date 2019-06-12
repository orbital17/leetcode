package leetcode

type RuneStack []rune

func (s *RuneStack) push(v rune) {
	*s = append(*s, v)
}

func (s *RuneStack) pop() (r rune, ok bool) {
	if len(*s) == 0 {
		return 0, false
	}

	r, ok = (*s)[len(*s)-1], true
	*s = (*s)[:len(*s)-1]
	return
}

func isValid020(s string) bool {
	stack := RuneStack{}
	for _, v := range s {
		switch v {
		case '(':
			stack.push(')')
		case '[':
			stack.push(']')
		case '{':
			stack.push('}')
		default:
			if last, ok := stack.pop(); !ok || last != v {
				return false
			}
		}
	}
	_, ok := stack.pop()
	return !ok
}
