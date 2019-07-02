package leetcode

import (
	"testing"

	. "./lib"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{"t", newListNode(1, 2), false},
		{"t", newListNode(1, 1), true},
		{"t", newListNode(1, 1, 2), false},
		{"t", newListNode(3, 1, 2, 1, 3), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
