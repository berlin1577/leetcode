package linked_list

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 超时，使用虚拟头节点，避免处理空节点
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{-1, nil}
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list2 != nil {
		tail.Next = list2
	} else if list1 != nil {
		tail.Next = list1
	}
	return head.Next
}
