package main

import (
    "fmt"
)

type TreeNode struct {
    value int
    left  *TreeNode
    right *TreeNode
}

func findLongestElement(root *TreeNode) int {
    if root == nil {
        return 0
    }
    // Traverse left and right subtrees
    leftDepth := findLongestElement(root.left)
    rightDepth := findLongestElement(root.right)

    // Return the maximum depth + 1 (current node)
    if leftDepth > rightDepth {
        return leftDepth + 1
    }
    return rightDepth + 1
}

func main() {
    root := &TreeNode{
        value: 10,
        left: &TreeNode{
            value: 5,
            left: &TreeNode{
                value: 3,
                left:  nil,
                right: nil,
            },
            right: &TreeNode{
                value: 8,
                left:  nil,
                right: nil,
            },
        },
        right: &TreeNode{
            value: 15,
            left: &TreeNode{
                value: 12,
                left:  nil,
                right: nil,
            },
            right: &TreeNode{
                value: 20,
                left:  nil,
                right: nil,
            },
        },
    }
    longest := findLongestElement(root)
    fmt.Println("The longest element in the tree is:", longest)
}
