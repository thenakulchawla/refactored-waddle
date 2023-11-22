package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createNewLinkedList(noOfNodes int) *ListNode {

	if noOfNodes <= 0 {
		return nil
	}

	dummyHead := &ListNode{}
	current := dummyHead

	for i := 1; i <= noOfNodes; i++ {
		current.Next = &ListNode{Val: i}
		current = current.Next

	}

	return dummyHead.Next

}

func printLinkedList(head *ListNode) {

	for head != nil {
		fmt.Println("val:", head.Val)
		head = head.Next
	}
}

func main() {

	head := createNewLinkedList(5)
	printLinkedList(head)

}
