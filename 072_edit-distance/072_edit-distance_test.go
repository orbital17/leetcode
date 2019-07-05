package leetcode

import "testing"

func Test_minDistance(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"basic", args{"horse", "ros"}, 3},
		{"basic", args{"", "ros"}, 3},
		{"basic", args{"", ""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
