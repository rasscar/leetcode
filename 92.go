func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	var dummy, prev, curr, next *ListNode
	dummy = &ListNode{0, head}
	prev = dummy

	var i int
	for i = 0; i < left-1; i++ {
		prev = prev.Next
	}

	curr = prev.Next

	for i = 0; i < right-left; i++ {
		next = curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}
