package arrays

import "fmt"

func ArrProduct() {
	nums := []int{1, 2, 3, 4}
	res := arrProduct(nums)
	fmt.Println(res)
}

func arrProduct(nums []int) []int {

	length := len(nums)

	l := make([]int, length)
	r := make([]int, length)
	ans := make([]int, length)

	l[0] = 1

	for i := 1; i < length; i++ {
		l[i] = nums[i-1] * l[i-1]
	}

	r[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		r[i] = nums[i+1] * r[i+1]
	}

	for i := 0; i < length; i++ {
		ans[i] = l[i] * r[i]
	}

	return ans
}
