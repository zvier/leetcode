package s2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSubArrayLen(t *testing.T) {
	type testcase struct {
		s    int
		ret  int
		nums []int
	}

	testcases := []testcase{
		{
			s:    7,
			ret:  2,
			nums: []int{2, 3, 1, 2, 4, 3},
		},
		{
			s:    15,
			ret:  5,
			nums: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tc := range testcases {
		ret := minSubArrayLen(tc.s, tc.nums)
		assert.Equal(t, tc.ret, ret, "they should be equal")
	}
}
