package leetcode

import "testing"
import . "../lib"

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
		{"basic", args{newListNode(1, 4, 5), 2}, newListNode(4, 1, 5)},
		{"basic", args{newListNode(1, 4, 5, 6), 3}, newListNode(5, 4, 1, 6)},
		{"basic", args{newListNode(1, 4, 5, 6, 7), 3}, newListNode(5, 4, 1, 6, 7)},
		{"basic", args{newListNode(1, 4, 5, 6), 5}, newListNode(1, 4, 5, 6)},
		{"basic", args{newListNode(1, 4, 5, 6, 7), 5}, newListNode(7, 6, 5, 4, 1)},
		{"basic", args{newListNode(1, 4, 5, 6, 7, 1, 4, 5, 6, 7), 5}, newListNode(7, 6, 5, 4, 1, 7, 6, 5, 4, 1)},
		{"basic", args{newListNode(1, 4, 5, 6, 7), 2}, newListNode(4, 1, 6, 5, 7)},
		{"basic", args{newListNode(1, 4, 5, 6, 7), 3}, newListNode(5, 4, 1, 6, 7)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); got.String() != tt.want.String() {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
