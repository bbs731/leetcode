package leetcode

import (
	"math"
)

/*

这种题目，搞死了我无数次， 打击了信心， 让我感觉自己是个傻子！

DP 的解法， 复杂度是 O（n^2)
但是还有数学的解法， 妈的！ O(1)

 */

 func twoEggDrop(n int) int {

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt64

		//  we try from 1st floor to ith floor
		// 然后我们选，这里面最小的
		for j := 1; j <= i; j++ {
			// suppose jth floor, egg no break , then its a subproblem
			ans := dp[i-j] + 1 // 1 try for no break

			// otherwise, we suppose jth floor egg break
			ans = max(ans, j-1-1+1+1)

			// 所有的 ans 里面选最小的
			dp[i] = min(dp[i], ans)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
