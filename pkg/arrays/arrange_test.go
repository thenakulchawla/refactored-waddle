package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoveZeroes(t *testing.T) {

	t.Run("base case", func(t *testing.T) {
		arr := []int{9, 7, 8, 0, 0, 8, 7, 6, 0, -1, 1}
		moveZeroes(arr)

		require.Equal(t, 0, arr[8])
		require.Equal(t, 0, arr[9])
		require.Equal(t, 0, arr[10])

	})

}

// TestMoveArrElements
func TestMoveArrElements(t *testing.T) {

	t.Run("base case", func(t *testing.T) {
		arr := []int{0, 2}
		moveArrElements(arr)

		require.Equal(t, arr[0], 0)
		require.Equal(t, arr[1], -1)
	})

	t.Run("longer case", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 8, 9}
		moveArrElements(arr)
		require.Equal(t, []int{-1, 1, 2, 3, 4, -1}, arr)
	})

	t.Run("empty case", func(t *testing.T) {
		arr := []int{}
		moveArrElements(arr)
		require.Equal(t, []int{}, arr)
	})

}

// TestMoveWithSwap
func TestMoveWithSwap(t *testing.T) {

	t.Run("base case", func(t *testing.T) {
		arr := []int{1, 0}
		moveArrElementsLessMemory(arr)

		require.Equal(t, arr[0], 0)
		require.Equal(t, arr[1], 1)
	})

	t.Run("longer case", func(t *testing.T) {
		arr := []int{9, 7, 6, 5, -1, -1, 2, 3, 8, -1}
		moveArrElementsLessMemory(arr)
		require.Equal(t, []int{-1, -1, 2, 3, -1, 5, 6, 7, 8, 9}, arr)
	})

	t.Run("empty case", func(t *testing.T) {
		arr := []int{}
		moveArrElementsLessMemory(arr)
		require.Equal(t, []int{}, arr)
	})

}
