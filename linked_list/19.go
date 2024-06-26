package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
*     Next *ListNode
 * }
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}
	slow, fast := dummy, head
	for i := 1; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
