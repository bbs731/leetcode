package leetcode

func rob2(nums []int) int {

	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for l := 2; l < n; l++ {
		dp[l] = max(dp[l-2]+nums[l], dp[l-1])
	}
	return dp[n-1]
}
