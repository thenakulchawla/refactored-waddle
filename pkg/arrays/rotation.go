package arrays

import "fmt"

// SearchInSortedRotated search in a rotated sorted numsay
// Leetcode 33: https://leetcode.com/problems/search-in-rotated-sorted-numsay/description/
func SearchInSortedRotated() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 6
	res := searchInSortedRotated(nums, target)
	fmt.Println(res)
}

// SearchSumInSortedRotated check if sum exists in a sorted rotated array
func SearchSumInSortedRotated() {
	nums := []int{5, 6, 7, 2, 3, 4}
	target := 11

	res := searchSumInSortedRotated(nums, target)
	fmt.Println(res)
}

// SearchInSortedRotatedNonDistinct search in non distinct rotated array
func SearchInSortedRotatedNonDistinct() {

	nums := []int{1, 0, 1, 1, 1}
	target := 0
	res := searchInSortedRotatedNonDistinct(nums, target)
	fmt.Println(res)
}

func findPivot(nums []int, low int, high int) int {

	if high < low {
		return -1
	}

	if high == low {
		return low
	}

	mid := (high + low) / 2

	if mid < high && nums[mid] > nums[mid+1] {
		return mid
	}

	if mid > low && nums[mid] < nums[mid-1] {
		return mid - 1
	}

	if nums[low] >= nums[mid] {
		return findPivot(nums, low, mid-1)
	}
	return findPivot(nums, mid+1, high)
}

func findPivotNonDisctint(nums []int, low int, high int) int {

	if high < low {
		return -1
	}

	if high == low {
		return low
	}

	mid := (high + low) / 2

	if mid < high && nums[mid] > nums[mid+1] {
		return mid
	}

	if mid > low && nums[mid] < nums[mid-1] {
		return mid - 1
	}

	if nums[low] > nums[mid] {
		return findPivot(nums, low, mid-1)
	}
	return findPivot(nums, mid+1, high)
}

func binarySearch(nums []int, low int, high int, targetber int) int {

	if high < low {
		return -1
	}

	mid := (high + low) / 2
	if nums[mid] == targetber {
		return mid
	}

	if nums[mid] < targetber {
		return binarySearch(nums, mid+1, high, targetber)
	}

	return binarySearch(nums, low, mid-1, targetber)
}

func searchInSortedRotated(nums []int, target int) int {
	n := len(nums)
	pivot := findPivot(nums, 0, n-1)

	if pivot == -1 {
		return binarySearch(nums, 0, n-1, target)
	}

	if nums[pivot] == target {
		return pivot
	} else if nums[0] <= target {
		return binarySearch(nums, 0, pivot-1, target)
	}

	return binarySearch(nums, pivot+1, n-1, target)

}

func searchSumInSortedRotated(nums []int, target int) bool {

	n := len(nums)
	pivot := findPivot(nums, 0, n-1)

	left := (pivot + 1) % n
	right := pivot

	for left != right {
		if nums[left]+nums[right] == target {
			return true
		}

		if nums[left]+nums[right] < target {
			left = (left + 1) % n
		} else {
			right = (n + right - 1) % n
		}
	}

	return false

}

func existsInFirst(nums []int, start int, target int) bool {
	return nums[start] <= target
}

func isBinarySearchHelpful(nums []int, left int, target int) bool {
	return nums[left] != target
}

func searchInSortedRotatedNonDistinct(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}

	start := 0
	end := n - 1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			return true
		}

		if !isBinarySearchHelpful(nums, start, nums[mid]) {
			start++
			continue
		}

		pivotArr := existsInFirst(nums, start, nums[mid])
		targetArr := existsInFirst(nums, start, target)

		if pivotArr != targetArr {

			if pivotArr {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if nums[mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return false
}
