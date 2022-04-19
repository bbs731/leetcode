package leetcode

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {

	// dp[n][p]  n 是人数， p 是 accum profit
	dp := make([][]int, n+1)
	l := len(profit)
	sum := 0
	// 感觉是 sum 太大的话，过不了 test case
	// 看看，官网的答案吧， 用的 是 minPorfit 作为上界
	for i := 0; i < l; i++ {
		sum += profit[i]
	}
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sum+1)
	}

	dp[0][0] = 1
	mod := 1000000007

	for i := 0; i < len(group); i++ {
		for j := n; j >= group[i]; j-- { // 因为是 0-1 背包的问题， 所以， j 和 k 都是从大往小推。  完全背包的问题，是从小往大推！
			for k := sum; k >= profit[i]; k-- {
				dp[j][k] += dp[j-group[i]][k-profit[i]]
				dp[j][k] %= mod
			}
		}
	}

	ans := 0
	for i := 0; i <= n; i++ {
		for j := minProfit; j <= sum; j++ {
			ans += dp[i][j]
			ans %= mod
		}
	}

	return ans
}
