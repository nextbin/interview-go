package easy

import "fmt"

func Run0237() {
	var n1 = ListNode{4, nil}
	var n2 = ListNode{5, &n1}
	var n3 = ListNode{1, &n2}
	var n4 = ListNode{9, &n3}
	deleteNode(&n2)
	fmt.Println(n4)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
