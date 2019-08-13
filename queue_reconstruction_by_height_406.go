package main

import (
	"fmt"
	"sort"
)

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(people))
}

// Queue is [][]int
type Queue [][]int

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q Queue) Less(i, j int) bool {
	p1, p2 := q[i], q[j]
	if p1[0] > p2[0] {
		return true
	} else if p1[0] == p2[0] {
		return p1[1] < p2[1]
	}
	return false
}

func reconstructQueue(people [][]int) [][]int {
	length := len(people)
	if length == 0 {
		return people
	}
	queue := make([][]int, length)
	sort.Sort(Queue(people))
	arrangePeople(people, length-1, queue)
	return queue
}

func arrangePeople(people [][]int, index int, queue [][]int) {
	person := people[index]
	if index == 0 {
		queue[0] = person
		return
	}
	arrangePeople(people, index-1, queue)
	k := person[1]
	copy(queue[k+1:], queue[k:])
	queue[k] = person
}

