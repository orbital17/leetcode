package leetcode

import "testing"

func Test_isValid020(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{"()"}, true},
		{args{"([])"}, true},
		{args{"([)]"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.args.s, func(t *testing.T) {
			if got := isValid020(tt.args.s); got != tt.want {
				t.Errorf("isValid020() = %v, want %v", got, tt.want)
			}
		})
	}
}
