package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateParenthesis(t *testing.T) {

	t.Run("case one", func(t *testing.T) {

		n := 3
		res := generateParenthesis(n)
		require.Equal(t, 5, len(res))

	})

}
