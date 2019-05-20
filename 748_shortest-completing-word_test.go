package leetcode

import "fmt"

func ExampleShortestCompletingWord() {
	list := []string{"step","steps","stripe","stepple"}
	res := shortestCompletingWord("1s3 PSt", list)
	fmt.Print(res)
	//Output:
	// steps
}

func ExampleShortestCompletingWord2() {
	list := []string{"looks", "pest", "stew", "show"}
	res := shortestCompletingWord("1s3 456", list)
	fmt.Print(res)
	//Output:
	// pest
}