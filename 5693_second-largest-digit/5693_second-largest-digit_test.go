package leetcode

import "testing"

func Test_secondHighest(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{
				"dfa12321afd",
			},
			2,
		},
		{
			args{
				"dfa11111afd",
			},
			-1,
		},
		{
			args{
				"dfa11101afd",
			},
			0,
		},
		{
			args{
				"dfaafd",
			},
			-1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run("tt.name", func(t *testing.T) {
			if got := secondHighest(tt.args.s); got != tt.want {
				t.Errorf("secondHighest() = %v, want %v", got, tt.want)
			}
		})
	}
}
