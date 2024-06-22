package linkedlist

type SNode struct {
	Val  int
	Next *SNode
}

type SingleLL struct {
	head *SNode
	tail *SNode
	len  int
}

type DNode struct {
	Val  int
	Next *SNode
	prev *SNode
}

type DoubleLL struct {
	len  int
	head *DNode
	tail *DNode
}
