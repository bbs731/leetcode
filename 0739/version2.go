package leetcode


/*
这种答案， 看起来就是对的！
 */
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	counts := make([]int, n)
	// stack 用来放一个， 递减的 单调栈。（里面的内容放的是对应的 temperatures 里的 Index)
	stack := []int{0}

	for i := 1; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			counts[top] = i - top
			// pop stack
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return counts
}
