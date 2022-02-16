package leetcode


/*
这个复杂度是， O（n ^2)  有  O（n sqrt(n)) 的解法 ！
 */
func minSteps(n int) int {

	dp := make([]int, n+1)
	// 初始化
	for i := 0; i <= n; i++ {
		dp[i] = i
	}

	dp[1] = 0

	for i := 3; i <= n; i++ {
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				dp[i] = min(dp[i], dp[j]+1+i/j-1)
			}
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
