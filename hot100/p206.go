package hot100

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	// prev 是转换好的节点
	// next临时变量
	current := head
	
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
