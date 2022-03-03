package leetcode

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	dp := make([]int, len(s))

	// 预处理， dp[i] 表示， index = i 的 ( 对应的 ) 的 index 位置
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			if len(stack) == 0 {
				// invalid bracket
				continue
			}
			index := stack[len(stack)-1] // top '(' index
			dp[index] = i                // record the pair
			stack = stack[:len(stack)-1]
		} else {
			// otherwise we have '('
			stack = append(stack, i)
		}
	}
	ans := 0

	for i := 0; i < len(s); i++ {
		length := 0
		for i < len(s) && dp[i] != 0 {
			length += dp[i] - i + 1
			i = dp[i] + 1
		}
		ans = max(ans, length)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
