package hot100

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
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head.Next

	for fast.Next != nil && fast.Next.Next != nil {
		if slow.Val == fast.Val {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
