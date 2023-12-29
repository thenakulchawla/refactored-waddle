package arrays

import "fmt"

func CombinationSum() {

	nums := []int{2, 3, 6, 7}
	target := 7
	ans := combinationSum(nums, target)
	fmt.Println(ans)
}

func combinationSum(candidates []int, target int) [][]int {

	var result [][]int
	backtrack(target, make([]int, 0), 0, candidates, &result)
	return result
}

func backtrack(remain int, comb []int, start int, candidates []int, res *[][]int) {

	if remain == 0 {
		copyComb := make([]int, len(comb))
		copy(copyComb, comb)
		*res = append(*res, copyComb)
	} else if remain < 0 || start > len(candidates) {
		return
	}

	for i := start; i < len(candidates); i++ {
		comb = append(comb, candidates[i])
		backtrack(remain-candidates[i], comb, i, candidates, res)
		comb = comb[:len(comb)-1]
	}

}
