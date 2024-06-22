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
	backtrackCS(target, make([]int, 0), 0, candidates, &result)
	return result
}

func backtrackCS(remain int, comb []int, start int, candidates []int, res *[][]int) {

	if remain == 0 {
		copyComb := make([]int, len(comb))
		copy(copyComb, comb)
		*res = append(*res, copyComb)
	} else if remain < 0 || start > len(candidates) {
		return
	}

	for i := start; i < len(candidates); i++ {
		comb = append(comb, candidates[i])
		backtrackCS(remain-candidates[i], comb, i, candidates, res)
		comb = comb[:len(comb)-1]
	}

}

func GenParenthesis() {
	n := 3
	res := generateParenthesis(n)
	fmt.Println(res)
}

func generateParenthesis(n int) []string {
	var res []string
	backtrackGP(&res, n, 0, 0, "")
	return res
}

func backtrackGP(res *[]string, n, open, close int, currSequence string) {

	if open == close && close == n {
		*res = append(*res, currSequence)
		return
	}

	if open < n {
		currSequence += "("
		open++
		backtrackGP(res, n, open, close, currSequence)
		open--
		currSequence = currSequence[:len(currSequence)-1]
	}

	if close < open {
		currSequence += ")"
		close++
		backtrackGP(res, n, open, close, currSequence)
		close--
		currSequence = currSequence[:len(currSequence)-1]
	}

}
