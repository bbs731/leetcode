package leetcode

/*

https://leetcode-cn.com/leetbook/read/dynamic-programming-1-plus/5audnm/
官网上，给出的状态转移方程：

dp[i][j] = dp[i + 1][j] + 1;
dp[i][j] = min(dp[i][j], dp[i + 1][k] + dp[k + 1][j]); 其中 i < k <= j 且 s[i] == s[k]

 */

func strangePrinter(s string) int {

	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	getter := func(i, j int) int {
		if i > j {
			return 0
		}
		return dp[i][j]
	}
	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			dp[i][j] = dp[i+1][j] + 1

			for k := i + 1; k <= j; k++ {
				if s[i] == s[k] {
					dp[i][j] = min(dp[i][j], dp[i+1][k]+getter(k+1, j))
				}
			}
		}
	}
	return dp[0][n-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
