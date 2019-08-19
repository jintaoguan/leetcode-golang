package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head1 := &ListNode{
		Val: 3,
	}
	intersection := &ListNode{
		Val: 2,
	}
	head1.Next = intersection
	intersection.Next = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  4,
			Next: intersection,
		},
	}
	head2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	fmt.Println(detectCycle(head1).Val)
	fmt.Println(detectCycle(head2))
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1, p2 := head, head
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
		if p2 != nil {
			p2 = p2.Next
		}
		if p2 == nil || p2.Next == nil {
			return nil
		}
		if p1 == p2 {
			break
		}
	}
	p1 = head
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

