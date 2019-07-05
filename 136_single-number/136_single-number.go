package leetcode

func singleNumber(nums []int) int {
	var a uint
	for _, v := range nums {
		a ^= uint(v)
	}
	return int(a)
}
