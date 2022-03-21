package leetcode

/*
就这点玩仍，你写了一个小时？

 */
func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	// dp 来表示， dp[i] 用前 ith 个数，包括 ith 组成的数组， 满足条件的子数组的个数。  dp[i] 就是一个递增的数组。
	dp := make([]int, n+1)

	if nums[0] >= k {
		dp[0] = 0
	} else {
		dp[0] = 1
	}

	startwindow := 0
	product := nums[0]

	/*
	下面是基于双指针的想法，来统计。 官网上的解， 更巧妙的处理了  statwindow == i 的情况
	 */

	for i := 1; i < n; i++ {
		product = product * nums[i]
		if product < k {
			dp[i] = dp[i-1] + i - startwindow + 1
		} else {
			for product >= k && startwindow < i {
				product = product / nums[startwindow]
				startwindow++
			}
			// special case
			if startwindow == i {
				// product = 1
				if nums[i] >= k {
					dp[i] = dp[i-1]
				} else {
					dp[i] = dp[i-1] + 1
				}
				product = nums[i]
			} else {
				dp[i] = dp[i-1] + i - startwindow + 1
			}
		}
	}
	return dp[n-1]
}
