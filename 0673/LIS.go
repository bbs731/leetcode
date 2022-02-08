package leetcode
/*
这道题，对我来说，有点难！ 
 */
func findNumberOfLIS(nums []int) int {

	dp := make([]int, len(nums)) // dp[i] = max{dp[j] + 1}
	cnt := make([]int, len(nums))

	dp[0] = 1
	cnt[0] = 1
	maxlen := 0
	ans := 0

	for i := 1; i < len(nums); i++ {
		cnt[i] = 1
		dp[i] = 1
		for j := 0; j <= i-1; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1 // take the max
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}

		if dp[i] > maxlen {
			maxlen = dp[i]
			ans = cnt[i]
		} else if dp[i] == maxlen {
			ans += cnt[i]
		}
	}
	return ans
}
