package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	dict := map[string]int{}
	result := [][]string{}
	for _, v := range strs {
		bytes := []byte(v)
		sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
		key := string(bytes)
		if index, ok := dict[key]; ok {
			result[index] = append(result[index], v)
		} else {
			dict[key] = len(result)
			result = append(result, []string{v})
		}
	}
	return result
}
