package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}

//func partition(head *ListNode, x int) *ListNode {
//	if head == nil && head.Next == nil {
//		return nil
//	}
//	t := &ListNode{-1, head}
//	left, right := t, t
//	for right.Next.Next != nil {
//		if right.Next.Val >= x {
//			right = right.Next
//		} else {
//			p := right.Next
//			right.Next = p.Next
//			p.Next = left.Next
//			left.Next = p
//		}
//		right = right.Next
//	}
//	return t.Next
//}
