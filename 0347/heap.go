package leetcode

import "container/heap"

type MyInt [][2]int

func (h MyInt) Len() int {
	return len(h)
}

func (h MyInt) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h MyInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MyInt) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *MyInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// 先计数， 然后再 partition

	table := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		table[nums[i]]++
	}

	freqs := &MyInt{}

	heap.Init(freqs)
	for k, v := range table {
		heap.Push(freqs, [2]int{k, v})
	}

	count := len(table) - k
	for count > 0 {
		count--
		heap.Pop(freqs)
	}

	ans := []int{}
	for freqs.Len() != 0 {
		t := freqs.Pop().([2]int)
		ans = append(ans, t[0])
	}
	return ans
}
