package leetcode

/*
这是一个错误的答案， 看看问题出在哪儿里！
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
			dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			for k := i; k < j; k++ {
				if s[k] == s[j] {
					dp[i][j] = max(dp[i][j], dp[k+1][j-1]+2)
					break
				}
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
