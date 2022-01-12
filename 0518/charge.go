package leetcode

/*
这也太牛了， 一次就过！
 */
func change(amount int, coins []int) int {

	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = 0 // 0 表示，没有组合的方案
	}
	dp[0] = 1
	// 一次，只考虑增加一种货币的来。 【1，2，5],
	// 先考虑 只用1块的 update dp[i] 然后再考虑，使用2块的，来更新 dp[i], 以此类推
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] > 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}
