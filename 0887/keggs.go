package leetcode

import "math"

func superEggDrop(k int, n int) int {

	// dp[n][k] 代表的是， n 层楼， k 个 eggs, 最少能得到 f 的次数
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}

	// 初始化
	for eggs := 1; eggs <= k; eggs++ {
		dp[1][eggs] = 1
	}

	for i :=1; i<=n; i++{
		dp[i][1] = i
	}

	for i := 2; i <= n; i++ {
		// eggs 从 2 开始， eggs 等于 1 是 dp[i][1] 是在初始化中完成的， 在 dp[j-1][eggs-1] 有 eggs- 1操作， 因为 eggs=0 dp[i][0] 没有明确定义， 所以让 eggs 从2 开始
		for eggs := 2; eggs <= k; eggs++ {
			dp[i][eggs] = math.MaxInt64
			for j := 1; j <= i; j++ {
				// suppose jth floor , eggs no break.
				ans := dp[i-j][eggs] + 1

				// otherwise jth floor, eggs break
				ans = max(ans, dp[j-1][eggs-1]+1)

				dp[i][eggs] = min(dp[i][eggs], ans)

			}
		}
	}
	return dp[n][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
