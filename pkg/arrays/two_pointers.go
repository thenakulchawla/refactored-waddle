package arrays

// maxArea : container with most water
// LC11: https://leetcode.com/problems/container-with-most-water/description/
func maxArea(height []int) int {

	mArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		width := right - left
		hl := height[left]
		hr := height[right]

		area := width * min(hl, hr)
		mArea = max(mArea, area)

		if hl <= hr {
			left++
		} else {
			right--
		}

	}

	return mArea

}
