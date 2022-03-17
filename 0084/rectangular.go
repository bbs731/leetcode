package leetcode

/*
官网的思路，更巧妙。

1. 不需要记录 times.  把 index 放进单调栈里面。 这样根据 index 和 i 的值，可以计算长方形的宽了！
2. 用特殊的 前后 哨兵 0 值， 不需要特殊处理 0 了， 同时最后的 stack 清空也可以不做了。
 */
func largestRectangleArea(heights []int) int {
	// 需要一个 递增的单调栈

	type item struct {
		val    int
		counts int
	}
	stack := make([]item, 0)

	stack = append(stack, item{heights[0], 1})
	ans := heights[0]

	for i := 1; i < len(heights); i++ {
		sl := len(stack)
		if sl != 0 {
			if heights[i] >= stack[sl-1].val {
				// insert
				stack = append(stack, item{heights[i], 1})
			} else {
				times := 1
				// pop until we can insert
				for sl > 0 && heights[i] < stack[sl-1].val {
					// pop
					lastitem := stack[sl-1]
					ans = max(ans, lastitem.val*lastitem.counts)
					stack = stack[:sl-1]
					sl--
					if sl == 0 {
						times += lastitem.counts
					} else {
						// accumulate counts  // 需要比较， 是加在 stack top 的item 还是 heights[i] 上
						if heights[i] >= stack[sl-1].val {
							times += lastitem.counts
						} else {
							stack[sl-1].counts += lastitem.counts
						}
					}
				}
				// insert
				// times handle the case the stack is empty and insert the ith item
				if heights[i] != 0 {
					stack = append(stack, item{heights[i], times})
				}
			}
		} else {
			if heights[i] != 0 {
				stack = append(stack, item{heights[i], 1})
			}
		}
	}
	// at the end, if stack is not empty, handle that again
	l := len(stack)
	counts := 0
	for l > 0 {
		counts += stack[l-1].counts
		ans = max(ans, stack[l-1].val*counts)
		// pop
		l--
	}
	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
