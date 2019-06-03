package leetcode

func trap(hs []int) int {
	if len(hs) < 3 {
		return 0
	}
	l, r := 0, len(hs)-1
	lMax, rMax := hs[l], hs[r]
	totalArea := 0
	for r > l+1 {
		if lMax < rMax {
			l++
			if hs[l] > lMax {
				lMax = hs[l]
			} else {
				totalArea += lMax - hs[l]
			}
		} else {
			r--
			if hs[r] > rMax {
				rMax = hs[r]
			} else {
				totalArea += rMax - hs[r]
			}
		}
	}
	return totalArea
}
