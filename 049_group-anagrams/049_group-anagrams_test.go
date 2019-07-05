package leetcode

import (
	"fmt"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{"basic", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, "[[eat tea ate] [tan nat] [bat]]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.strs); fmt.Sprint(got) != tt.want {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
