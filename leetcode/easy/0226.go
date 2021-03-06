package easy

import "fmt"

func Run0226() {
	fmt.Println(invertTree(nil))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	var tmp = root.Left
	root.Left = root.Right
	root.Right = tmp
	return root
}
