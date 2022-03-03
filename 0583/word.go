package leetcode

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	// 把初始化的过程，放到了 getter 的逻辑里面。 这样难以理解啊，高超的技巧，做对了，只是走运啊！
	var getter = func(i, j int) int {
		if i == -1 && j == -1 {
			return 0
		}
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		return dp[i][j]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = getter(i-1, j-1)
			} else {
				dp[i][j] = min(getter(i-1, j), getter(i, j-1)) + 1
			}
		}
	}

	return dp[n-1][m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
