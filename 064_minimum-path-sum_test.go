package leetcode

import "fmt"

func Example_minPathSum() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	res := minPathSum(grid)
	fmt.Println(res)
	//Output:
	//7
}
