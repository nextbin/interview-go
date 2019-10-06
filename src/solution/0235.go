package solution

import "fmt"

func Run0235() {
	var a = TreeNode{1, nil, nil}
	var b = TreeNode{2, &a, nil}
	fmt.Println(lowestCommonAncestor(&b, &a, &b))
	a = TreeNode{2, nil, nil}
	b = TreeNode{1, nil, &a}
	fmt.Println(lowestCommonAncestor(&b, &a, &b))
}

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val > q.Val {
		var tmp = p
		p = q
		q = tmp
	}
	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	}
	var ret *TreeNode
	ret = lowestCommonAncestor(root.Left, p, q)
	if ret != nil {
		return ret
	}
	ret = lowestCommonAncestor(root.Right, p, q)
	if ret != nil {
		return ret
	}
	return nil
}
