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
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	head2 := &ListNode{
		Val:  2,
		Next: head1.Next.Next,
	}
	fmt.Println(getIntersectionNode1(head1, head2).Val)
	fmt.Println(getIntersectionNode2(head1, head2).Val)
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	counterA, counterB := 0, 0
	curA, curB := headA, headB
	for curA != nil {
		counterA++
		curA = curA.Next
	}
	for curB != nil {
		counterB++
		curB = curB.Next
	}
	diff := diff(counterA, counterB)
	var curL, curS *ListNode
	if counterA > counterB {
		curL = headA
		curS = headB
	} else {
		curL = headB
		curS = headA
	}
	for diff != 0 {
		diff--
		curL = curL.Next
	}
	for curL != curS {
		curL = curL.Next
		curS = curS.Next
	}
	return curL
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)
	curA, curB := headA, headB
	for curA != nil {
		visited[curA] = true
		curA = curA.Next
	}
	for curB != nil {
		if _, ok := visited[curB]; ok {
			return curB
		}
		curB = curB.Next
	}
	return nil
}

