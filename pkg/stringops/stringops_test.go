package stringops

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalindromeWithReverese(t *testing.T) {

	t.Run("base case", func(t *testing.T) {

		s := "A man, a plan, a canal: Panama"
		res := palindromeWithReverse(s)
		require.True(t, res)
	})

}

func TestPalindromeWithPointers(t *testing.T) {

	t.Run("base case", func(t *testing.T) {

		s := "A man, a plan, a canal: Panama"
		res := palindromeWithTwoPointer(s)
		require.True(t, res)
	})

}

func TestReverseVowels(t *testing.T) {

	t.Run("base case", func(t *testing.T) {

		s := "hello"
		res := reverseVowels(s)
		require.Equal(t, "holle", res)
	})

}

func TestReverseWords(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		s := "the sky is blue"
		rev := reverseWords(s)
		require.Equal(t, "blue is sky the", rev)
	})

	t.Run("with spaces", func(t *testing.T) {
		s := "  hello world  "
		rev := reverseWords(s)
		require.Equal(t, "world hello", rev)

	})

	t.Run("with internal space", func(t *testing.T) {
		s := "a good   example"
		rev := reverseWords(s)
		require.Equal(t, "example good a", rev)

	})

}

func TestIsSubsequence(t *testing.T) {

	t.Run("true case", func(t *testing.T) {
		require.True(t, isSubsequence("abc", "ahbgdc"))
	})

	t.Run("false case", func(t *testing.T) {
		require.True(t, isSubsequence("axc", "ahbgdc"))
	})

}

func TestCloseStrings(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		ans := closeStrings("a", "aa")
		require.False(t, ans)
	})

	t.Run("case 2", func(t *testing.T) {
		ans := closeStrings("cabbba", "abbccc")
		require.False(t, ans)
	})

	t.Run("case 3", func(t *testing.T) {
		ans := closeStrings("cabbbaaaaa", "abbcccbbbb")
		require.False(t, ans)
	})

}
