package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func BuildBSTFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Value: preorder[0]}
	i := 1

	for i < len(preorder) && preorder[i] < root.Value {
		i++
	}

	root.Left = BuildBSTFromPreorder(preorder[1:i])
	root.Right = BuildBSTFromPreorder(preorder[i:])

	return root
}

func InorderTraversal(root *TreeNode) {
	if root != nil {
		InorderTraversal(root.Left)
		fmt.Printf("%d ", root.Value)
		InorderTraversal(root.Right)
	}
}

func main() {
	preorder := []int{8, 5, 1, 7, 10, 12}

	root := BuildBSTFromPreorder(preorder)

	fmt.Println("In-order traversal of the constructed BST:")
	InorderTraversal(root)
	fmt.Println()
}
// Traversal Order:
// In the first code, the inorder traversal is used, which prints the elements of the BST in sorted order.
// In the second code, the inorder traversal is also used, so it prints the elements in sorted order as well.
// Traversal Sequence Difference:
// In the first code, the elements are inserted into the BST based on a postorder traversal sequence, which means the root is the last element in the sequence.
// In the second code, the elements are inserted into the BST based on a preorder traversal sequence, which means the root is the first element in the sequence.