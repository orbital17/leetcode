package leetcode

var mapping = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func makeCombination(digits string, n int) string {
	res := make([]byte, len(digits))
	p := 1
	for i := len(digits) - 1; i >= 0; i-- {
		v := digits[i]
		mi := n / p % len(mapping[v])
		p *= len(mapping[v])
		res[i] = mapping[v][mi]
	}
	return string(res)
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	n := 1
	for i := range digits {
		v := digits[i]
		n *= len(mapping[v])
	}
	result := make([]string, n)
	for i := range result {
		result[i] = makeCombination(digits, i)
	}
	return result
}
