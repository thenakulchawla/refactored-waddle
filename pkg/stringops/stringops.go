package stringops

import (
	"fmt"
	"strings"
	"unicode"
)

func Palindrome() {
	s := "A man, a plan, a canal: Panama"
	res := palindromeWithTwoPointer(s)
	fmt.Println(res)
}

func palindromeWithReverse(s string) bool {

	sRunes := []rune(s)
	sFilteredRunes := []rune{}

	for _, r := range sRunes {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sFilteredRunes = append(sFilteredRunes, unicode.ToLower(r))
		}
	}

	sReverseRunes := []rune{}
	for i := len(sFilteredRunes) - 1; i >= 0; i-- {
		sReverseRunes = append(sReverseRunes, unicode.ToLower(sFilteredRunes[i]))
	}

	originalString := string(sFilteredRunes)
	fmt.Println(originalString)
	reverseString := string(sReverseRunes)
	fmt.Println(reverseString)

	return originalString == reverseString

}

func palindromeWithTwoPointer(s string) bool {

	f := func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return unicode.ToLower(r)
		}

		return -1
	}

	s = strings.Map(f, s)

	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i += 1
		j -= 1
	}

	return true

}

func reverseVowels(s string) string {

	vowelMap := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	var vowelArr []rune

	runeArr := []rune(s)

	for i := len(runeArr) - 1; i >= 0; i-- {
		if _, ok := vowelMap[runeArr[i]]; ok {
			vowelArr = append(vowelArr, runeArr[i])
		}
	}

	count := 0
	for i := 0; i < len(runeArr); i++ {
		if _, ok := vowelMap[runeArr[i]]; ok {
			runeArr[i] = vowelArr[count]
			count++
			if count == len(vowelArr) {
				return string(runeArr)
			}
		}
	}

	return string(runeArr)

}

func reverseWords(s string) string {

	s = strings.TrimSpace(s)
	words := strings.Fields(s)

	low, high := 0, len(words)-1

	for low < high {
		words[low], words[high] = words[high], words[low]
		low++
		high--
	}

	rev := strings.Join(words, " ")

	return rev

}

func isSubsequence(s string, t string) bool {

	sRune := []rune(s)
	tRune := []rune(t)
	leftBound := len(sRune)
	rightBound := len(tRune)

	left, right := 0, 0

	for left < leftBound && right < rightBound {
		if sRune[left] == tRune[right] {
			left++
		}
		right++
	}

	return left == leftBound
}

func closeStrings(word1 string, word2 string) bool {

	length := len(word1)
	if len(word1) != len(word2) {
		return false
	}

	chars1 := make([]int, 26)
	chars2 := make([]int, 26)

	values1 := make([]int, length+1)
	values2 := make([]int, length+1)

	for _, c := range word1 {
		index := c - 'a'
		chars1[index]++
		values1[chars1[index]]++
	}

	for _, c := range word2 {
		chars2[c-'a']++
		values2[chars2[c-'a']]++
	}

	for i := 0; i < 26; i++ {
		if chars1[i] > 0 && chars2[i] == 0 {
			return false
		}
	}

	for i := 0; i < length; i++ {
		if values1[i] != values2[i] {
			return false
		}
	}

	return true

}
