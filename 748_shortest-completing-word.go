package leetcode

import (
	"strings"
)

func shortestCompletingWordArr(licensePlate string, words []string) string {
	l := letters(strings.ToLower(licensePlate))
	minLen := -1
	res := ""
	for _, w := range words {
		if (minLen < 0 || len(w) < minLen) && check(l, w) {
			res = w
			minLen = len(w)
		}
	}
	return res
}

func letters(word string) []int {
	l := make([]int, 26)
	for _, v := range word {
		i := int(v) - int('a')
		if i >= 0 && i < 26 {
			l[i]++
		}
	}
	return l
}

func check(l []int, word string) bool {
	wordL := letters(word)
	for i := 0; i < 26; i++ {
		if wordL[i] < l[i] {
			return false
		}
	}
	return true
}

func shortestCompletingWord(licensePlate string, words []string) string {
	l := primeProduct(strings.ToLower(licensePlate))
	minLen := -1
	res := ""
	for _, w := range words {
		if (minLen < 0 || len(w) < minLen) && primeProduct(w)%l == 0 {
			res = w
			minLen = len(w)
		}
	}
	return res
}

var primes [26]int = [26]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

func primeProduct(word string) int {
	r := 1
	for _, v := range strings.ToLower(word) {
		i := int(v) - int('a')
		if i >= 0 && i < 26 {
			r *= primes[i]
		}
	}
	return r
}
