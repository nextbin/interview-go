package easy

import "fmt"

func Run0206() {
	var n1 = ListNode{1, nil}
	var n2 = ListNode{2, &n1}
	var n3 = ListNode{6, &n2}
	var n4 = ListNode{6, &n3}
	var n5 = ListNode{7, &n4}
	var res = reverseList(&n5)
	fmt.Println(res)
	n1 = ListNode{1, nil}
	n2 = ListNode{2, &n1}
	n3 = ListNode{6, &n2}
	n4 = ListNode{6, &n3}
	res = reverseList(&n4)
	fmt.Println(res)
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre = head
	var cur = head.Next
	pre.Next = nil
	for cur != nil {
		var nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
