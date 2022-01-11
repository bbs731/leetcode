package leetcode

/*
Time complexit is O(n^2)
 */
func longestPalindromeSubseq(s string) int {
	N := len(s)

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		dp[i][i] = 1
	}

	for length := 2; length <= N; length++ {
		for i := 0; i+length-1 < N; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][N-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
