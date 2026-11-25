func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var tail *ListNode
	tail = head
	for i := 0; i < k; i++ {
		if tail == nil {
			return head
		}
		tail = tail.Next
	}

	var newHead *ListNode
	newHead = reverse(head, tail)
	head.Next = reverseKGroup(tail, k)
	return newHead
}

func reverse(cur *ListNode, end *ListNode) *ListNode {
	var prev *ListNode
	prev = nil

	for cur != end {
		var nxt *ListNode
		nxt = cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	return prev
}
