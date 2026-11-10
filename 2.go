func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy, tail, node *ListNode
	dummy = &ListNode{Val: 0}
	tail = dummy

	var carry, d1, d2, jum, digit int
	carry = 0

	for l1 != nil || l2 != nil || carry != 0 {
		d1 = 0
		if l1 != nil {
			d1 = l1.Val
		}

		d2 = 0
		if l2 != nil {
			d2 = l2.Val
		}

		jum = d1 + d2 + carry
		digit = jum % 10
		carry = jum / 10

		node = &ListNode{Val: digit}
		tail.Next = node
		tail = tail.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return dummy.Next
}
