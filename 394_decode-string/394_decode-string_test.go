package leetcode

import "testing"

func Test_decodeString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"t", "3[a]2[bc]", "aaabcbc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
