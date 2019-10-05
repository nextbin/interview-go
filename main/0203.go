package main

import "fmt"

func main() {
	var n1 = ListNode{1, nil}
	var n2 = ListNode{2, &n1}
	var n3 = ListNode{6, &n2}
	var n4 = ListNode{6, &n3}
	var n5 = ListNode{7, &n4}
	var res = removeElements(&n5, 6)
	fmt.Println(res)
	n1 = ListNode{1, nil}
	n2 = ListNode{2, &n1}
	n3 = ListNode{6, &n2}
	n4 = ListNode{6, &n3}
	res = removeElements(&n4, 6)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	var ret = head
	for ret != nil {
		if ret.Val == val {
			ret = ret.Next
		} else {
			break
		}
	}
	if ret != nil {
		var pre = ret
		var cur = pre.Next
		for cur != nil {
			if cur.Val == val {
				cur = cur.Next
				pre.Next = cur
			} else {
				pre = cur
				cur = cur.Next
			}
		}
	}
	return ret
}
