package leetcode

func minimumTotal(triangle [][]int) int {
	N := len(triangle)
	if N == 1 {
		return triangle[0][0]
	}

	pdp := []int{triangle[0][0]}

	for i := 1; i < N; i++ {
		len := len(triangle[i])
		dp := make([]int, len)
		// dp[0] 和 dp[len-1] 这两是边界条件
		dp[0] = pdp[0] + triangle[i][0]
		dp[len-1] = pdp[len-2] + triangle[i][len-1]
		for j := 1; j < len-1; j++ {
			dp[j] = min(pdp[j], pdp[j-1]) + triangle[i][j]
		}
		pdp = dp
	}

	ans := pdp[0]
	for i := 1; i < len(triangle[N-1]); i++ {
		ans = min(ans, pdp[i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
