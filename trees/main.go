package main

import (
	"fmt"

	stack "github.com/golang-collections/collections/stack"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	var res []int
	// stack := stack.New[*TreeNode]()
	stack := stack.New()

	current := root

	for current != nil || stack.Len() > 0 {
		for current != nil {
			stack.Push(current)
			current = current.Left
		}

		current = stack.Pop().(*TreeNode)
		if current == nil {
			break
		}

		res = append(res, current.Val)
		current = current.Right

	}

	return res

}

func createSpecificTree() *TreeNode {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}

	// Manually link nodes
	root.Right = node2
	root.Left = node3

	return root

}
func createGenerictree(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}

	i := 1

	for i < len(values) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root

}

func main() {
	input := createSpecificTree()
	res := inorderTraversal(input)

	for _, val := range res {
		fmt.Println(val)

	}

}
