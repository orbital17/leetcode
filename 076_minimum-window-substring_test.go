package leetcode

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"T", args{"ADOBECODEBANC", "ABC"}, "BANC"},
		{"T", args{"A", "A"}, "A"},
		{"T", args{"aa", "aa"}, "aa"},
		{"T", args{"Acc", "A"}, "A"},
		{"T", args{"ccA", "A"}, "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
