package leetcode


/*
DP problem 873 的进阶版本。
本题，没要求， 数组，是完全升序的，所以，数组中的数字，可以重复出现
 */
func longestArithSeqLength(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	index := make(map[int][]int)
	for i := 0; i < n; i++ {
		index[nums[i]] = append(index[nums[i]], i)
	}

	ans := 0
	for i := 2; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			diff := nums[i] - nums[j]
			numsk := nums[j] - diff
			if ks, ok := index[numsk]; ok {
				for _, k := range ks {
					if k < j {
						dp[i][j] = dp[j][k] + 1
						ans = max(ans, dp[i][j])
					}
				}
			}
		}
	}

	return ans + 2
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
