package leetcode

func countBits(num int) []int {
	res := make([]int, num+1)
	cur := 1
	for cur < num+1 {

		for i := 0; i < cur && i+cur < num+1; i++ {
			res[i+cur] = res[i] + 1
		}
		cur *= 2
	}
	return res
}
