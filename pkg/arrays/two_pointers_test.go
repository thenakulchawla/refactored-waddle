package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxArea(t *testing.T) {

	t.Run("base case", func(t *testing.T) {
		height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		res := maxArea(height)
		require.Equal(t, 49, res)
	})
}
