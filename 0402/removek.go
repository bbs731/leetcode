package leetcode

import "strings"
/*
这是，你有史以来，做的，最拉跨的一道题， 你做了一个半小时，就是一道 medium
 */
func removeKdigits(num string, k int) string {

	if len(num) == 0 {
		return "0"
	}
	stack := make([]byte, 0)

	for i := 0; i < len(num); i++ {
		if len(stack) == 0 {
			stack = append(stack, num[i])
			continue
		}
		top := stack[len(stack)-1]
		if num[i] >= top {
			stack = append(stack, num[i])
			continue
		}
		// otherwise we need to pop the stack
		for num[i] < top && k > 0 {
			// pop the stack
			stack = stack[:len(stack)-1]
			k--
			if len(stack) > 0 {
				top = stack[len(stack)-1]
			} else {
				break
			}
		}
		// add num[i] to the top of the stack
		stack = append(stack, num[i])
	}

	for k > 0 {
		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
	}

	ans := string(stack)
	ans = strings.TrimLeft(ans, "0")
	if ans == "" {
		return "0"
	}
	return ans
}
