package medium

import (
	"fmt"
	"math"
)

func Run0098() {
	var n1 = TreeNode{1, nil, nil}
	var n3 = TreeNode{3, nil, nil}
	var n2 = TreeNode{2, &n1, &n3}
	fmt.Println(isValidBST(&n2))
	n1 = TreeNode{1, nil, nil}
	n3 = TreeNode{3, nil, nil}
	var n6 = TreeNode{6, nil, nil}
	var n4 = TreeNode{4, &n3, &n6}
	var n5 = TreeNode{5, &n1, &n4}
	fmt.Println(isValidBST(&n5))
	fmt.Println(isValidBST(&n4))
	fmt.Println(isValidBST(nil))
	n1 = TreeNode{1, nil, nil}
	n2 = TreeNode{1, nil, &n1}
	fmt.Println(isValidBST(&n2))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var isValidLeft = true
	if root.Left != nil {
		isValidLeft = isValidBSTHelper(root.Left, math.MinInt64, root.Val)
	}
	var isValidRight = true
	if root.Right != nil && isValidLeft {
		isValidRight = isValidBSTHelper(root.Right, root.Val, math.MaxInt64)
	}
	return isValidLeft && isValidRight
}

func isValidBSTHelper(root *TreeNode, min int, max int) bool {
	if root.Val <= min || root.Val >= max {
		return false
	}
	var isValidLeft = true
	if root.Left != nil {
		isValidLeft = isValidBSTHelper(root.Left, min, root.Val)
	}
	var isValidRight = true
	if root.Right != nil && isValidLeft {
		isValidRight = isValidBSTHelper(root.Right, root.Val, max)
	}
	return isValidLeft && isValidRight
}
