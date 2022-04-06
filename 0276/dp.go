package leetcode

func numWays(n int, k int) int {

	if n <= 1 {
		return k
	}

	dp := make([]int, n) // dp 表示  i 和 i-1 的颜色不同的数量
	cp := make([]int, n) // cp 表示  i 和 i -1 颜色相同的数量

	/*
		这个初始化，是懵的啊！
	 */
	dp[0] = k
	cp[0] = 0

	for i := 1; i < n; i++ {
		cp[i] = dp[i-1]
		dp[i] = (k-1)*dp[i-1] + (k-1)*cp[i-1]
	}

	return dp[n-1] + cp[n-1]

}
