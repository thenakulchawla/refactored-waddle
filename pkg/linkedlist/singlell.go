package linkedlist

func (s *SingleLL) insertAtBack(element int) {

	newNode := &SNode{Val: element}

	if s.head == nil {
		s.head = newNode
		return
	}

	current := s.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (s *SingleLL) insertBeforeFirst(element int) {
	newNode := &SNode{Val: element}

	if s.head == nil {
		s.head = newNode
		return
	}

	newNode.Next = s.head
	s.head = newNode
}

func (s *SingleLL) insertAfterNodeValue(element int, nodeVal int) bool {

	newNode := &SNode{Val: element}

	curr := s.head

	for curr != nil {
		if curr.Val == nodeVal {
			newNode.Next = curr.Next
			curr.Next = newNode
			return true
		}

		curr = curr.Next
	}

	return false

}

func (s *SingleLL) reverseIterate() {

	if s.head == nil || s.head.Next == nil {
		return
	}

	var prev *SNode
	curr := s.head

	for curr != nil {

		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	s.head = prev

}

func (s *SingleLL) reverseRecurse() {

	if s.head == nil || s.head.Next == nil {
		return
	}

	curr := s.head
	rev := recursiveReverse(curr)
	// printUsingHead(rev)
	s.head = rev

}

func recursiveReverse(node *SNode) *SNode {

	if node == nil || node.Next == nil {
		return node
	}

	p := recursiveReverse(node.Next)
	node.Next.Next = node
	node.Next = nil
	return p

}
