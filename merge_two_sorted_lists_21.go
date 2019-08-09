package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	head2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
		},
	}
	fmt.Println(mergeTwoLists(head1, head2).Val)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	head := &ListNode{
		Val: 0,
	}
	last := head
	ptr1, ptr2 := l1, l2
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val < ptr2.Val {
			last.Next = ptr1
			last = ptr1
			ptr1 = ptr1.Next
		} else {
			last.Next = ptr2
			last = ptr2
			ptr2 = ptr2.Next
		}
	}
	if ptr1 != nil {
		last.Next = ptr1
	}
	if ptr2 != nil {
		last.Next = ptr2
	}
	return head.Next
}

