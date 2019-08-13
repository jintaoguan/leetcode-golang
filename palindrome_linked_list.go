package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
		},
	}
	fmt.Println(isPalindrome1(head))
	fmt.Println(isPalindrome2(head))
}

func isPalindrome1(head *ListNode) bool {
	stack := make([]int, 0)
	cur := head
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		if head.Val != stack[i] {
			return false
		}
		head = head.Next
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	tail1, p1, p2 := head, head, head
	for p2 != nil {
		tail1 = p1
		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			break
		}
		p2 = p2.Next
	}
	tail1.Next = nil
	head2 := reverseList(p1)
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	newHead := reverseList(head.Next)
	next.Next = head
	head.Next = nil
	return newHead
}

