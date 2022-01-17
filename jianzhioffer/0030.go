package offer

/*
you are perfect, baby!
 */

// 用空间换时间呗，保存两个栈
type MinStack struct {
	stack    []int
	minstack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minstack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)

	if len(this.minstack) == 0 || this.minstack[len(this.minstack)-1] >= x {
		this.minstack = append(this.minstack, x)
	}

}

func (this *MinStack) Pop() {

	if len(this.stack) == 0 {
		return
	}
	elem := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if elem == this.minstack[len(this.minstack)-1] {
		this.minstack = this.minstack[:len(this.minstack)-1]
	}

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) Min() int {
	return this.minstack[len(this.minstack)-1]

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
