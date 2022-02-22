package leetcode

/*

经典的 DP 问题，  LCS 的变种吗？

dp[i][j] =   if nums1[i] == nums[j]    dp[i-1][j-1] + 1
			else   max (dp[i-1][j], dp[i][j-1] )
 */

func maxUncrossedLines(nums1 []int, nums2 []int) int {

	n1 := len(nums1)
	n2 := len(nums2)
	dp := make([][]int, n1)

	var getter = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		return dp[i][j]
	}

	for i := 0; i < n1; i++ {
		dp[i] = make([]int, n2)
		for j := 0; j < n2; j++ {

			if nums1[i] == nums2[j] {
				dp[i][j] = getter(i-1, j-1) + 1
			} else {
				dp[i][j] = max(getter(i-1, j), getter(i, j-1))
			}

		}
	}

	return dp[n1-1][n2-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
