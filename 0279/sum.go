package leetcode

/*
 你也太厉害了吧， 一次救过！
 */
func numSquares(n int) int {
	nums := make([]int, 0)

	for i := 1; i*i <= n; i++ {
		nums = append(nums, i)
	}

	// candidate nums
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		// 1 永远可以组成任何 number 对吧？
		dp[i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j]*nums[j] > i {
				break
			}
			dp[i] = min(dp[i], dp[i-nums[j]*nums[j]]+1)

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
