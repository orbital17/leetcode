package main

import (
	"./lib"
	"testing"
)

func Test(t *testing.T) {
	err := lib.NewProblem("287_find-the-duplicate-number")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
