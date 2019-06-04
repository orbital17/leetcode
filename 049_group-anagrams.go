package leetcode

func groupAnagrams(strs []string) [][]string {
	dict := map[[26]uint8]int{}
	result := [][]string{}
	for _, v := range strs {
		key := [26]uint8{}
		for _, symbol := range v {
			key[symbol-'a']++
		}
		if index, ok := dict[key]; ok {
			result[index] = append(result[index], v)
		} else {
			dict[key] = len(result)
			result = append(result, []string{v})
		}
	}
	return result
}
