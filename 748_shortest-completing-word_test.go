package leetcode

import "fmt"

func Example_shortestCompletingWord() {
	list := []string{"step", "steps", "stripe", "stepple"}
	res := shortestCompletingWord("1s3 PSt", list)
	fmt.Print(res)
	//Output:
	// steps
}

func Example_shortestCompletingWord2() {
	list := []string{"looks", "pest", "stew", "show"}
	res := shortestCompletingWord("1s3 456", list)
	fmt.Print(res)
	//Output:
	// pest
}
