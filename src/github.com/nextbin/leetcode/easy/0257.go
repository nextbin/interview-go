package easy

import "fmt"

func Run0257() {
	fmt.Println(binaryTreePaths(nil))
	var a = TreeNode{1, nil, nil}
	var b = TreeNode{2, &a, nil}
	fmt.Println(binaryTreePaths(&b))
	fmt.Println(binaryTreePaths(&a))
	var n5 = TreeNode{5, nil, nil}
	var n3 = TreeNode{3, nil, nil}
	var n2 = TreeNode{2, nil, &n5}
	var n1 = TreeNode{1, &n2, &n3}
	fmt.Println(binaryTreePaths(&n1))

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}
	var ret = []string{}
	if root.Left != nil {
		var res1 = binaryTreePaths(root.Left)
		for _, s := range res1 {
			ret = append(ret, fmt.Sprintf("%d->%s", root.Val, s))
		}
	}
	if root.Right != nil {
		var res1 = binaryTreePaths(root.Right)
		for _, s := range res1 {
			ret = append(ret, fmt.Sprintf("%d->%s", root.Val, s))
		}
	}
	return ret
}
