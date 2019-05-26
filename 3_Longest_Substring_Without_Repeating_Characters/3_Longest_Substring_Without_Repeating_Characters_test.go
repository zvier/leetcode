package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		In  string
		Out int
	}{
		{
			In:  "abcabcbb",
			Out: 3,
		},
		{
			In:  "bbbbb",
			Out: 1,
		},
		{
			In:  "pwwkew",
			Out: 3,
		},
	}
	for _, cas := range cases {
		in := cas.In
		expectOut := cas.Out
		out := lengthOfLongestSubstring(in)
		if expectOut != out {
			t.Errorf("%s the length of in's longest substring expect %d, but real output is %d", in, expectOut, out)
		}
	}
}
