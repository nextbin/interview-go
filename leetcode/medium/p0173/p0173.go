package p0173

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	vals []int
	idx  int
}

func Constructor(root *TreeNode) BSTIterator {
	it := &BSTIterator{}
	it.handle(root)
	return *it
}

func (b *BSTIterator) Next() int {
	val := b.vals[b.idx]
	b.idx++
	return val
}

func (b *BSTIterator) HasNext() bool {
	return b.idx < len(b.vals)
}

func (b *BSTIterator) handle(root *TreeNode) {
	if root.Left != nil {
		b.handle(root.Left)
	}
	b.vals = append(b.vals, root.Val)
	if root.Right != nil {
		b.handle(root.Right)
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
