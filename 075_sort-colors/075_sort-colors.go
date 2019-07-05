package leetcode

func sortColors(nums []int) {
	a, b := 0, 0
	for i, v := range nums {
		switch v {
		case 1:
			if b != i {
				nums[b] = 1
				nums[i] = 2
			}
			b++
		case 0:
			if a != i {
				nums[a] = 0
				if b != a {
					nums[b] = 1
				}
				if b != i {
					nums[i] = 2
				}
			}
			a++
			b++
		}
	}
}
