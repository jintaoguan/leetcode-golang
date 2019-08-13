package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{3, 0, 1, 0}
	k := 1
	fmt.Println(topKFrequent(nums, k))
}

// An TopKNode is the node in TopKHeap
type TopKNode struct {
	Val  int
	Freq int
}

// An TopKHeap is the heap to track frequency
type TopKHeap []*TopKNode

func (h TopKHeap) Len() int {
	return len(h)
}
func (h TopKHeap) Less(i, j int) bool {
	return h[i].Freq >= h[j].Freq
}
func (h TopKHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push an element to heap
func (h *TopKHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*TopKNode))
}

// Pop element from heap
func (h *TopKHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[0 : l-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, v := range nums {
		if _, ok := counts[v]; ok {
			counts[v]++
		} else {
			counts[v] = 1
		}
	}
	h := make(TopKHeap, len(counts))
	index := 0
	for key, value := range counts {
		h[index] = &TopKNode{Val: key, Freq: value}
		index++
	}
	result := make([]int, k)
	heap.Init(&h)
	for i := 0; i < k; i++ {
		element := heap.Pop(&h).(*TopKNode).Val
		result[i] = element
	}
	return result
}

