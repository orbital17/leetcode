package main

import (
	"./lib"
	"fmt"
	"os"
)

func main() {
	err := lib.NewProblem(os.Args[1])
	if err != nil {
		fmt.Printf("error: %v", err)
	}
}
