package leetcode

import "container/heap"
/*
这是一个噩梦啊！
 */

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *IntHeap) Pop() interface{} {
	// pop the last elements
	old := *h
	n := len(old)
	t := old[n-1]
	*h = old[0 : n-1]
	return t
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func getLeastNumbers(arr []int, k int) []int {

	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < len(arr); i++ {
		heap.Push(h, arr[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]int, 0)
	for h.Len() > 0 {
		ans = append(ans, h.Pop().(int))
	}
	return ans
}
