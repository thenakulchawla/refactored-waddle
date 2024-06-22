package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidParenthesis(t *testing.T) {

	t.Run("base case", func(t *testing.T) {
		s := "()"
		res := isValid(s)
		require.True(t, res)
	})

}
