package leetcode

import (
	"testing"

	. "github.com/orbital17/leetcode/lib"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"basic", args{NewListNode(1, 4, 5), 2}, NewListNode(4, 1, 5)},
		{"basic", args{NewListNode(1, 4, 5, 6), 3}, NewListNode(5, 4, 1, 6)},
		{"basic", args{NewListNode(1, 4, 5, 6, 7), 3}, NewListNode(5, 4, 1, 6, 7)},
		{"basic", args{NewListNode(1, 4, 5, 6), 5}, NewListNode(1, 4, 5, 6)},
		{"basic", args{NewListNode(1, 4, 5, 6, 7), 5}, NewListNode(7, 6, 5, 4, 1)},
		{"basic", args{NewListNode(1, 4, 5, 6, 7, 1, 4, 5, 6, 7), 5}, NewListNode(7, 6, 5, 4, 1, 7, 6, 5, 4, 1)},
		{"basic", args{NewListNode(1, 4, 5, 6, 7), 2}, NewListNode(4, 1, 6, 5, 7)},
		{"basic", args{NewListNode(1, 4, 5, 6, 7), 3}, NewListNode(5, 4, 1, 6, 7)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); got.String() != tt.want.String() {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
