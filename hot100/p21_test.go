package hot100

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ans := &ListNode{
		Val:  0,
		Next: nil,
	}

	head := ans

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			ans.Next = list2
			list2 = list2.Next
		} else {
			ans.Next = list1
			list1 = list1.Next
		}
		ans = ans.Next
	}

	if list1 != nil {
		ans.Next = list1
	}

	if list2 != nil {
		ans.Next = list2
	}

	return head.Next
}
