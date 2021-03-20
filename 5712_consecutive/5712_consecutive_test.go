package leetcode

import "testing"

func Test_getMaximumConsecutive(t *testing.T) {
	type args struct {
		coins []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{[]int{1, 3}},
			2,
		},
		{
			args{[]int{1, 1, 1, 4}},
			8,
		},
		{
			args{[]int{1, 4, 10, 3, 1}},
			20,
		},
	}
	for _, tt := range tests {
		t.Run("tt.name", func(t *testing.T) {
			if got := getMaximumConsecutive(tt.args.coins); got != tt.want {
				t.Errorf("getMaximumConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
