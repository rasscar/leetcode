func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	ekor := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			ekor.Next = l1
			l1 = l1.Next
		} else {
			ekor.Next = l2
			l2 = l2.Next
		}
		ekor = ekor.Next
	}
	if l1 != nil {
		ekor.Next = l1
	}
	if l2 != nil {
		ekor.Next = l2
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	hasil := lists[0]
	for i := 1; i < len(lists); i++ {
		hasil = mergeTwoLists(hasil, lists[i])
	}
	return hasil
}
