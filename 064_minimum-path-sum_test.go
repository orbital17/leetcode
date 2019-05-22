package leetcode

import "fmt"

func ExampleMinPathSum() {
	grid := [][]int{
		[]int{1,3,1},
		[]int{1,5,1},
		[]int{4,2,1},
	  }
	res := minPathSum(grid)
	fmt.Println(res)
	//Output:
	//7
}