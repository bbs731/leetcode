package leetcode

/*
边界条件很容易出错啊！
 */
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)

	for _, s := range wordDict {
		dict[s] = true
	}

	N := len(s)
	dp := make([][]bool, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]bool, N)
	}

	// 初始化
	for i := 0; i < N; i++ {
		//dp[i][i] = true
		for j := i; j < N; j++ {
			if _, ok := dict[s[i:j+1]]; ok {
				dp[i][j] = true
			}
		}
	}

	for len := 1; len <= N; len++ {
		for i := 0; i < N; i++ {
			j := i + len
			if j < N && dp[i][j] == true {
				continue
			}
			for k := i; j < N && k < j; k++ {
				if dp[i][k] && dp[k+1][j] {
					dp[i][j] = true
					break
				}
			}
		}
	}
	return dp[0][N-1]
}
