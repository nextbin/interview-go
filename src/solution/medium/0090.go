package medium

func Run0090() {
	var n5 = ListNode{5, nil}
	var n4 = ListNode{4, &n5}
	var n3 = ListNode{3, &n4}
	var n2 = ListNode{2, &n3}
	var n1 = ListNode{1, &n2}
	reverseBetween(&n1, 2, 4).PrintList()
	n5 = ListNode{5, nil}
	reverseBetween(&n5, 1, 1).PrintList()
	n5 = ListNode{5, nil}
	n4 = ListNode{4, &n5}
	n3 = ListNode{3, &n4}
	n2 = ListNode{2, &n3}
	n1 = ListNode{1, &n2}
	reverseBetween(&n1, 1, 5).PrintList()
	n5 = ListNode{5, nil}
	n4 = ListNode{4, &n5}
	n3 = ListNode{3, &n4}
	n2 = ListNode{2, &n3}
	n1 = ListNode{1, &n2}
	reverseBetween(&n1, 1, 3).PrintList()
	n5 = ListNode{5, nil}
	n4 = ListNode{4, &n5}
	n3 = ListNode{3, &n4}
	n2 = ListNode{2, &n3}
	n1 = ListNode{1, &n2}
	reverseBetween(&n1, 3, 5).PrintList()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	var before *ListNode = nil
	var after *ListNode = nil
	var pre *ListNode = nil
	var cur = head
	var cnt = 1
	var isReverse = false
	var reverseHead *ListNode = nil
	var reverseTail *ListNode = nil
	for cur != nil {
		if cnt == m {
			before = pre
			isReverse = true
			reverseHead = cur
		}
		if isReverse {
			var nxt = cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		} else {
			pre = cur
			cur = cur.Next
		}
		cnt++
		if cnt == n+1 {
			isReverse = false
			reverseTail = pre
			after = cur
		}
	}
	if before != nil {
		before.Next = reverseTail
	}
	reverseHead.Next = after
	if before == nil {
		return reverseTail
	}
	return head

}
