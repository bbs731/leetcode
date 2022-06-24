package leetcode

/*
it is not that hard!
 */
func winnerSquareGame(n int) bool {
	dp := make([]bool, n+1)
	dp[0] = false

	candidates := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		dp[i*i] = true
		candidates = append(candidates, i*i)
	}

	for i := 2; i <= n; i++ {
		for _, c := range candidates {
			if i-c >= 0 {
				if dp[i-c] == false {
					dp[i] = true
				}
			} else {
				break
			}
		}
	}
	return dp[n]
}