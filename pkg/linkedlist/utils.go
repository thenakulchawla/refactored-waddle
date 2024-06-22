package linkedlist

import "github.com/rs/zerolog/log"

// Print function to call from cobra directly
func Print() {

	ll := createSLL(5)
	ll.print()

}

func createSLL(len int) *SingleLL {

	dummy := &SNode{}
	localHead := dummy

	currLength := 0
	for i := 0; i < len; i++ {
		nodeToAdd := &SNode{Val: i}
		localHead.Next = nodeToAdd
		localHead = localHead.Next
		currLength += 1
	}

	return &SingleLL{
		head: dummy.Next,
		len:  currLength,
	}

}

func (s *SingleLL) length() int {
	currLength := 0
	curr := s.head

	for curr != nil {
		currLength += 1
		curr = curr.Next
	}

	return currLength

}

func (s *SingleLL) lastElement() *SNode {
	curr := s.head
	for curr.Next != nil {
		curr = curr.Next
	}

	return curr
}

func (s *SingleLL) print() []int {
	var res []int

	curr := s.head

	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}

	log.Info().Ints("result", res).Msg("print single linked list")

	return res
}

func printUsingHead(head *SNode) []int {

	var res []int

	curr := head

	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}

	log.Info().Ints("result", res).Msg("sll using head")
	return res

}
