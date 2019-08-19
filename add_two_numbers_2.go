package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	fmt.Println(addTwoNumbers(l1, l2).Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	p1, p2 := l1, l2
	dummy := &ListNode{Val: 0}
	pre := dummy
	for p1 != nil || p2 != nil || carry != 0 {
		sum := 0
		if p1 != nil {
			sum += p1.Val
		}
		if p2 != nil {
			sum += p2.Val
		}
		if carry != 0 {
			sum += carry
		}
		digit := sum % 10
		carry = sum / 10
		newNode := &ListNode{Val: digit}
		pre.Next = newNode
		pre = newNode
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	return dummy.Next
}

