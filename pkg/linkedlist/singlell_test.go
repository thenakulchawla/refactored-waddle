package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrintSingleLL(t *testing.T) {

	t.Run("base case", func(t *testing.T) {

		ll := createSLL(5)

		require.Equal(t, 5, ll.len)
		require.Equal(t, 4, ll.length())
		ll.print()

	})

}

func TestSLLInsertion(t *testing.T) {

	t.Run("insert at back", func(t *testing.T) {

		sll := createSLL(5)

		sll.insertAtBack(100)
		backEl := sll.lastElement()
		require.Equal(t, backEl.Val, 100)
		sll.print()

	})

	t.Run("insert before first", func(t *testing.T) {
		sll := createSLL(5)

		sll.insertBeforeFirst(100)
		backEl := sll.lastElement()
		require.Equal(t, backEl.Val, 100)
		sll.print()

	})

	t.Run("insert after node value", func(t *testing.T) {
		sll := createSLL(5)
		nodeVal := 3
		e := 8

		sll.insertAfterNodeValue(e, nodeVal)
		require.Equal(t, sll.length(), 6)
		sll.insertAfterNodeValue(e, 4)
		require.Equal(t, sll.length(), 7)
		sll.print()

	})

}

func TestReverse(t *testing.T) {

	t.Run("iterative", func(t *testing.T) {

		sll := createSLL(5)
		sll.reverseIterate()

		revRes := sll.print()
		require.Equal(t, []int{4, 3, 2, 1, 0}, revRes)
	})

	t.Run("recursive", func(t *testing.T) {

		sll := createSLL(5)
		sll.reverseRecurse()

		revRes := sll.print()
		require.Equal(t, []int{4, 3, 2, 1, 0}, revRes)
	})

}
