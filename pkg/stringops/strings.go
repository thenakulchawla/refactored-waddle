package stringops

import (
	"fmt"
	"unicode"
)

func Palindrome() {
	s := "A man, a plan, a canal: Panama"
	res := palindromeWithReverse(s)
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
