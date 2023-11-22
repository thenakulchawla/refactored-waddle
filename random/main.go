package main

import (
	"fmt"
)

func mid() {

	left := 0
	right := 9

	pivot := (left + right) / 2
	fmt.Println(pivot)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindrome(s string) string {
	start := 0
	end := 0

	for i := range s {
		len1 := expand_from_middle(s, i, i)
		len2 := expand_from_middle(s, i, i+1)

		max_len := 0
		if len1 > len2 {
			max_len = len1
		} else {
			max_len = len2
		}

		if max_len > end-start {
			start = i - ((max_len - 1) / 2)
			end = i + (max_len / 2)
		}

	}

	return s[start : end+1]

}

func longestPalindrome1(s string) string {
	start := 0
	end := 0

	for i := range s {
		len_odd := expand_from_middle(s, i, i)
		len_even := expand_from_middle(s, i, i+1)

		max_len := max(len_odd, len_even)

		if max_len > end-start {
			start = i - ((max_len - 1) / 2)
			end = i + (max_len / 2)
		}

	}

	return s[start : end+1]
}

func expand_from_middle(s string, left int, right int) int {

	if s == "" || left > right {
		return 0
	}

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}

	return right - left - 1
}

func lengthOfLongestSubstring(s string) int {

	hash_map := make(map[rune]int)

	res := 0
	i := 0
	for j, c := range s {
		if k, ok := hash_map[c]; ok {
			i = max(k, i)
		}
		res = max(res, j-i+1)
		hash_map[c] = j + 1
	}

	return res

}

func find_item(list []int, item int) *int {

	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return &mid
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	// for something
	return nil

}

func for_loop() {

	a := "nakul"

	for i, c := range a {
		fmt.Println("index: ", i)
		fmt.Println("character: ", c)
	}
}

func main() {

	// result := longestPalindrome1("babad")
	// fmt.Println(result)
	list := []int{1, 2, 3, 4, 5}

	result := find_item(list, 2)
	fmt.Println(*result)

}
