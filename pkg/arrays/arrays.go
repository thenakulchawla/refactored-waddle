package arrays

import (
	"fmt"
)

func LongCons() {
	nums := []int{9, 8, 7, -1, 0, 2, 3, 1}
	res := longestConsecutive(nums)
	fmt.Println(res)
}

func longestConsecutive(nums []int) int {

	res := 1
	numsMap := make(map[int]bool)

	for _, num := range nums {
		numsMap[num] = true
	}

	for num := range numsMap {
		if ok, _ := numsMap[num-1]; !ok {
			currNum := num
			currStreak := 1

			for numsMap[currNum+1] {
				currNum += 1
				currStreak += 1
			}

			res = max(res, currStreak)

		}

	}

	return res

}
