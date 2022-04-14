package leetcode

/*
多年以前， 在 USACO 上见过的题目！
dfs  + cache  = DP

还有， Loop 的形式， 尝试写一下！
 */
func PredictTheWinner(nums []int) bool {
	n := len(nums)
	prefixsum := make([]int, n+1)
	prefixsum[1] = nums[0]
	for i := 2; i <= n; i++ {
		prefixsum[i] = prefixsum[i-1] + nums[i-1]
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j ++ {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(start int, end int) int {
		if start == end {
			return nums[start]
		}
		if dp[start][end] != -1 {
			return dp[start][end]
		}

		total := prefixsum[end+1] - prefixsum[start]

		a := dfs(start+1, end)
		b := dfs(start, end-1)

		dp[start][end] = max(total-a, total-b)
		return dp[start][end]

	}

	score := dfs(0, n-1)
	if score >= prefixsum[n]-score {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
