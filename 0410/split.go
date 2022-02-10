package leetcode

import "math"

/*
DP problem

but also can be solved by greedy with binary search

这道题， 个人更倾向于  DP 的解法，  greedy 如意逻辑错误，不好证明啊！

 */
func splitArray(nums []int, m int) int {
	n := len(nums)
	// 准备一下前缀和
	sum := make([]int, n)
	sum[0] = nums[0]
	dp := make([][]int, n)
	dp[0] = make([]int, m+1)
	dp[0][1] = nums[0]
	maxnum := make([]int, n)
	maxnum[0] = nums[0]

	for i := 1; i < n; i++ {
		dp[i] = make([]int, m+1)
		sum[i] = sum[i-1] + nums[i]
		dp[i][1] = sum[i]
		maxnum[i] = max(maxnum[i-1], nums[i])
		if i < m {
			dp[i][i+1] = maxnum[i]
		}
	}
	if m == 1 {
		return sum[n-1]
	}

	for k := 2; k <= m; k++ {
		for i := k - 1; i < n; i++ {
			dp[i][k] = math.MaxInt64
			for j := i - 1; j >= k-2; j-- {
				dp[i][k] = min(dp[i][k], max(dp[j][k-1], sum[i]-sum[j]))
			}
		}
	}
	return dp[n-1][m]
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
