package leetcode

/*
	dp[i][target] = dp[i-1][target-1]* prob[i] + dp[i-1][target]*(1-prob[i])

这道题， 你根本就没想清楚！

可以降维吗？ target 的维度一定需要吗?
*/
func probabilityOfHeads(prob []float64, target int) float64 {
	n := len(prob)
	dp := make([][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]float64, target+1)
	}

	dp[0][0] = 1 - prob[0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] * (1 - prob[i])
	}
	var getter = func(index, k int) float64 {

		if index == -1 && k == 0 {
			return 1
		}
		if index == -1 {
			return 0
		}
		return dp[index][k]
	}
	for i := 0; i < n; i++ {
		for k := 1; k <= target && k <= i+1; k++ {
			dp[i][k] = getter(i-1, k-1)*prob[i] + getter(i-1, k)*(1-prob[i])
		}
	}

	return dp[n-1][target]
}
