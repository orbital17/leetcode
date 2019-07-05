package leetcode

// 1 2 3 4 5 6 7 8
//        ^
// m1 = 4
// res: 1 is to move m2 to the right
func checkMedian(nums1 []int, nums2 []int, m1 int, m2 int) int {
	if m1 < len(nums1) && m2 != 0 && nums1[m1] < nums2[m2-1] {
		return -1
	}
	if m2 < len(nums2) && m1 != 0 && nums2[m2] < nums1[m1-1] {
		return 1
	}
	return 0
}

func findMedian(nums1 []int, nums2 []int, m1 int, m2 int) float64 {
	var low, high int
	if m1 == 0 {
		low = nums2[m2-1]
	} else if m2 == 0 || nums1[m1-1] > nums2[m2-1] {
		low = nums1[m1-1]
	} else {
		low = nums2[m2-1]
	}
	if (len(nums1)+len(nums2))%2 == 1 {
		return float64(low)
	}
	if m1 == len(nums1) {
		high = nums2[m2]
	} else if m2 == len(nums2) || nums1[m1] < nums2[m2] {
		high = nums1[m1]
	} else {
		high = nums2[m2]
	}
	return float64(low+high) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	l1, l2 := len(nums1), len(nums2)
	m := (l1 + l2 + 1) / 2
	left, right := 0, l2
	var check, m2 int
	for {
		m2 = (left + right) / 2
		check = checkMedian(nums1, nums2, m-m2, m2)
		if check > 0 {
			if right-left == 1 {
				m2 = right
				break
			}
			left = m2
		}
		if check < 0 {
			right = m2
		}
		if check == 0 {
			break
		}
	}
	return findMedian(nums1, nums2, m-m2, m2)
}
