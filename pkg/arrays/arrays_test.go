package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestConsecutive(t *testing.T) {

	t.Run("base case", func(t *testing.T) {
		nums := []int{9, 8, 7, -1, 0, 2, 3, 1}
		res := longestConsecutive(nums)
		require.Equal(t, 5, res)
	})
}
