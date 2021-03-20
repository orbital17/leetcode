package leetcode

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func nearestValidPoint(x int, y int, points [][]int) int {
	var min, minIndex int
	minIndex = -1
	for i := 0; i < len(points); i++ {
		p := points[i]
		if p[0] == x || p[1] == y {
			d := distance(p, []int{x, y})
			if minIndex < 0 || d < min {
				minIndex = i
				min = d
			}
		}
	}
	return minIndex
}
