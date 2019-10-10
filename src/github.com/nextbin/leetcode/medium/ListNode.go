package medium

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) PrintList() {
	var cur = this
	if cur == nil {
		fmt.Println("nil")
		return
	}
	fmt.Print(cur.Val)
	cur = cur.Next
	for cur != nil {
		fmt.Print(fmt.Sprintf("->%d", cur.Val))
		cur = cur.Next
	}
	fmt.Println()
}
