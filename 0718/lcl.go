package leetcode

func findLength(nums1 []int, nums2 []int) int {

	// dp[i][j] 代表，以  ith, jth 结尾的连续的串的最大长度
	// dp[i][j] =  dp[i-1][j-1] + 1  // if nums[i] == nums[j]
	//			= 0  // if nums[i] != nums[j]
	N1 := len(nums1)
	N2 := len(nums2)
	dp := make([][]int, N1+1)
	for i := 0; i <= N1; i++ {
		dp[i] = make([]int, N2+1)
	}
	ans := 0
	for i := 1; i <= N1; i++ {
		for j := 1; j <= N2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				ans = max(ans, dp[i][j])
			} else {
				dp[i][j] = 0
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
