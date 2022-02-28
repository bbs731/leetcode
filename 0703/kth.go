package leetcode

/*
把 heap 的写法，多写，熟悉吧！ 
 */
import "container/heap"

type MyInts []int

func (h MyInts) Len() int {
	return len(h)
}

func (h MyInts) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MyInts) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *MyInts) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MyInts) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
	elements MyInts
	K        int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{MyInts{},
		k,
	}
	heap.Init(&kth.elements)
	for i := 0; i < len(nums); i++ {
		heap.Push(&kth.elements, nums[i])

		if kth.elements.Len() > k {
			heap.Pop(&kth.elements)
		}
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.elements, val)
	if this.elements.Len() > this.K {
		heap.Pop(&this.elements)
	}
	// return top
	return this.elements[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
