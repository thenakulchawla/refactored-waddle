package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindPivot(t *testing.T) {

	t.Run("increasing", func(t *testing.T) {
		arr := []int{0, 1, 2, 3, 4}
		n := findPivot(arr, 0, len(arr)-1)
		require.Equal(t, n, -1)
	})

	t.Run("rotated", func(t *testing.T) {
		arr := []int{5, 6, 7, 0, 1, 2}
		n := findPivot(arr, 0, len(arr)-1)
		require.Equal(t, n, 2)
		require.Equal(t, arr[n], 7)

	})

	t.Run("non distinct", func(t *testing.T) {
		nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}
		pivot := findPivotNonDisctint(nums, 0, len(nums)-1)
		require.Equal(t, 13, pivot)
	})

}

func TestBinarySearch(t *testing.T) {

	t.Run("increasing", func(t *testing.T) {
		arr := []int{0, 1, 2, 3, 5}
		n := searchInSortedRotated(arr, 2)
		require.Equal(t, n, 2)
		n1 := binarySearch(arr, 0, len(arr)-1, 5)
		require.Equal(t, n1, 4)

	})

	t.Run("rotated", func(t *testing.T) {
		arr := []int{5, 6, 7, 0, 1, 2}
		n := searchInSortedRotated(arr, 0)
		require.Equal(t, n, 3)
		require.Equal(t, arr[n], 0)
	})

	t.Run("edge case 1", func(t *testing.T) {
		arr := []int{1, 3}
		n := searchInSortedRotated(arr, 1)
		require.Equal(t, n, 0)
	})

	t.Run("edge case 2", func(t *testing.T) {
		arr := []int{1, 3, 5}
		n := searchInSortedRotated(arr, 1)
		require.Equal(t, n, 0)
	})

}

func TestSumInRotatedSorted(t *testing.T) {

	t.Run("base case", func(t *testing.T) {

		nums := []int{5, 6, 7, 2, 3, 4}
		target := 11

		check := searchSumInSortedRotated(nums, target)
		require.True(t, check)

	})
}

func TestInSortedRotetedNonDistint(t *testing.T) {

	t.Run("case1", func(t *testing.T) {
		nums := []int{1, 0, 1, 1, 1}
		target := 0
		res := searchInSortedRotatedNonDistinct(nums, target)
		require.True(t, res)
	})

	t.Run("case2", func(t *testing.T) {
		nums := []int{2, 5, 6, 0, 0, 1, 2}
		target := 0
		res := searchInSortedRotatedNonDistinct(nums, target)
		require.True(t, res)
	})

}
