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
	}
	fmt.Println(hasCycle(head))
	fmt.Println(hasCycle2(head))
}

func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	cur := head
	for true {
		if _, ok := visited[cur]; ok {
			return true
		}
		if cur == nil {
			return false
		}
		visited[cur] = true
		cur = cur.Next
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	p1, p2 := head, head
	for p2 != nil {
		p2 = p2.Next
		if p2 != nil {
			p2 = p2.Next
		} else {
			return false
		}
		p1 = p1.Next
		if p1 == p2 {
			return true
		}
	}
	return false
}

