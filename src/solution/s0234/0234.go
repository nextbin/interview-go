package s0234

import "fmt"

func Run0234() {
	var n1 = ListNode{2, nil}
	var n2 = ListNode{1, &n1}
	fmt.Println(isPalindrome(&n2))
	n1 = ListNode{1, nil}
	n2 = ListNode{2, &n1}
	var n3 = ListNode{2, &n2}
	var n4 = ListNode{1, &n3}
	fmt.Println(isPalindrome(&n4))
	fmt.Println(isPalindrome(nil))
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
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//1. 获取链表长度
	var size = 0
	var cur = head
	for cur != nil {
		size++
		cur = cur.Next
	}
	//2. 将链表均等切分为A、B（原链表长度为奇数时，将中间元素归属A）
	var tail1 = head
	var head2 = head
	var curSize = 0
	for curSize < (size+1)/2 {
		tail1 = head2
		head2 = head2.Next
		curSize++
	}
	tail1.Next = nil
	//3. 链表B反转
	head2 = reverseList(head2)
	//4. 比较链表A、B是否完全相等（原链表长度为奇数时，A比B多一个元素）
	var pt1 = head
	var pt2 = head2
	for pt2 != nil {
		if pt1.Val != pt2.Val {
			return false
		}
		pt1 = pt1.Next
		pt2 = pt2.Next
	}
	//5. 复原链表，避免需要外部使用
	head2 = reverseList(head2)
	tail1.Next = head2
	return true
}

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
