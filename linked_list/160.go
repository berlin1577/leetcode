package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	taila, tailb := headA, headB
	for taila != tailb {
		if taila == nil {
			taila = headB
		} else {
			taila = taila.Next
		}
		if tailb == nil {
			tailb = headA
		} else {
			tailb = tailb.Next
		}
	}
	return taila
}
