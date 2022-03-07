package leetcode

func largestSumOfAverages(nums []int, k int) float64 {
	// dp[i][k]  // [i, i+1，.... n-1] 分成 k 组 的最大平均值
	n := len(nums)
	dp := make([][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]float64, k+1)
	}

	//presum[i] 从 0 到 i 的sum 和
	presum := make([]int, n) // 前缀和， 用来计算 sum[i,j] = presum[j] - presum[i] + nums[i]
	presum[0] = nums[0]
	for i := 1; i < n; i++ {
		presum[i] = presum[i-1] + nums[i]
	}

	for i := n - 1; i >= 0; i-- {
		dp[i][1] = float64(presum[n-1]-presum[i]+nums[i]) / float64(n-i)
	}

	// 因为 k 只依赖 k-1 随意 dp[i][k] 可以降维成 dp[i]
	for p := 2; p <= k; p++ {
		for i := n - 1; i >= 0; i-- {
			avg := 0.0
			for j := i + 1; j < n; j++ {
				avg = max(avg, dp[j][p-1]+float64(presum[j-1]-presum[i]+nums[i])/float64(j-1-i+1))
			}
			dp[i][p] = avg
		}
	}
	return dp[0][k]
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
