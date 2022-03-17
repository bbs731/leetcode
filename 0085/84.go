package leetcode

func maximalRectangle(matrix [][]byte) int {

	m := len(matrix)    // number of rows
	n := len(matrix[0]) // cols

	ans := 0

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// first row
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
		}
	}
	ans = max(ans, largestRectangleArea(dp[0]))

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i-1][j] + 1
			}
		}
		ans = max(ans, largestRectangleArea(dp[i]))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	// 需要一个 递增的单调栈
	n := len(heights)
	stack := make([]int, 0)
	newheights := make([]int, 0, n+2)
	n = n + 2
	newheights = append(newheights, 0)
	newheights = append(newheights, heights...)
	newheights = append(newheights, 0)
	stack = append(stack, 0)
	ans := 0

	for i := 1; i < n; i++ {
		sl := len(stack)
		for sl > 0 && newheights[i] < newheights[stack[sl-1]] {
			// pop
			h := newheights[stack[sl-1]]
			rightindex := i
			//leftindex := stack[sl-1] leftindex 不能直接取， 因为在insert stack[sl-1] 的时候，可能 pop 了其他的element, 因此需要去，当前 stack top 的 index
			leftindex := stack[sl-2]

			w := rightindex - leftindex - 1
			ans = max(ans, h*w)
			stack = stack[:sl-1]
			sl--
		}
		// now insert element i
		stack = append(stack, i)
	}

	return ans
}
