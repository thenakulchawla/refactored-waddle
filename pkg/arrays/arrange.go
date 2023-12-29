package arrays

import (
	"fmt"

	"github.com/thenakulchawla/refactored-waddle/pkg/sorting"
)

// MoveZeroes: move all zeroes to the end
func MoveZeroes() {
	arr := []int{9, 7, 8, 0, 0, 8, 7, 6, 0, -1, 1}
	moveZeroes(arr)
	fmt.Println(arr)
}

/*
CheckIndex if element for index i exists, place it or put -1
if input is [1,0], return [0,1]
if input is [2,0], return [0,-1]
*/
func MoveArrElements() {

	arr := []int{1, 0}
	moveArrElements(arr)
	fmt.Println(arr)

}

func moveZeroes(arr []int) {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}

	for count < len(arr) {
		arr[count] = 0
		count++
	}

}

func moveArrElements(arr []int) {

	numMap := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		numMap[arr[i]] = true
	}

	for i := 0; i < len(arr); i++ {
		if ok, _ := numMap[i]; ok {
			arr[i] = i
		} else {
			arr[i] = -1
		}
	}
}

func moveArrElementsLessMemory(arr []int) {

	i := 0
	for i < len(arr) {
		if arr[i] > 0 && arr[i] != i {
			arr[arr[i]], arr[i] = arr[i], arr[arr[i]]

		} else {
			i += 1
		}
	}

}

// GroupAnagrams group the anagrams together
// LC49: https://leetcode.com/problems/group-anagrams/description/
func GroupAnagrams() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)

}

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	strMap := make(map[string][]string)

	for _, s := range strs {
		sortedString := sorting.SortString(s)
		strMap[sortedString] = append(strMap[sortedString], s)
	}

	for _, v := range strMap {

		var tmpList []string

		for _, anagram := range v {
			tmpList = append(tmpList, anagram)
		}

		res = append(res, tmpList)
	}

	return res
}
