package arrays

import "fmt"

func ValidParenthesis() {

	s := "()"
	res := isValid(s)
	fmt.Println(res)
}

func isValid(s string) bool {

	n := len(s)
	if n%2 != 0 {
		return false
	}

	stack := []rune{}

	open := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	for _, rn := range s {
		if closer, ok := open[rn]; ok {
			stack = append(stack, closer)
			continue
		}

		currLength := len(stack) - 1
		if currLength < 0 || stack[currLength] != rn {
			return false
		}

		stack = stack[:currLength]
	}

	return len(stack) == 0

}
