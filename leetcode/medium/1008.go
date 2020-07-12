package medium


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := TreeNode{preorder[0],nil,nil}
	for i, val := range preorder {
		if i == 0 {
			continue
		}
		insertVal(&root, val)
	}
	return &root
}

func insertVal(root *TreeNode, val int) {
	if root.Val >= val && root.Left == nil {
		root.Left = &TreeNode{val, nil,nil}
	} else if root.Val < val && root.Right == nil {
		root.Right = &TreeNode{val, nil,nil}
	}else if root.Val > val {
		insertVal(root.Left, val)
	}else {
		insertVal(root.Right,val)
	}
}