package leetcode

import "testing"
import "fmt"


func printPermutation(nums []int) {
	for _,v := range nums{
		fmt.Print(v)
	}
	fmt.Println("")
}

func Test_nextPermutation(t *testing.T) {
	nums := []int{4, 3, 2, 1}
	for i:=0; i<24; i++ {
		nextPermutation(nums)
		printPermutation(nums)
	}
}
