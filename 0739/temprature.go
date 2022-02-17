package leetcode

/*
看看 version2.go 。 统计 cnt++ 太繁琐了， 其实可以直接用 index 的差值获得， 反而因为 cnt, 结果还产生了一些边界条件
 */
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	counts := make([]int, n)
	// stack 用来放一个， 递减的 单调栈。（里面的内容放的是对应的 temperatures 里的 Index)
	stack := []int{0}

	for i := 1; i < n; i++ {
		cnt := 0
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if temperatures[i] > temperatures[top] {
				cnt++
				counts[top] += cnt
				// pop stack
				stack = stack[:len(stack)-1]
			} else {
				// accumulate
				for k := 0; k < len(stack); k++ {
					counts[stack[k]] += cnt
				}
				break
			}
		}
		stack = append(stack, i)
	}
	// 还要有个特殊处理。 如果最后 Stack 不为空的话，需要 mark 剩下所有的为0
	for i := 0; i < len(stack); i++ {
		counts[stack[i]] = 0
	}

	return counts
}
