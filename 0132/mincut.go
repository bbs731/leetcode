package leetcode

func minCut(s string) int {

	n := len(s)
	// p[i][j] 表示字符串 s[i:j+1] 是否是 palidrome
	p := make([][]bool, n)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		p[i] = make([]bool, n)
		dp[i] = make([]int, n)
	}

	// 初始化 p  with len 0 and 1
	for i := 0; i < n; i++ {
		p[i][i] = true
		j := i + 1
		if j < n && s[i] == s[j] {
			p[i][j] = true
		}
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			j := i + l
			if j < n {
				if s[i] == s[j] && p[i+1][j-1] {
					p[i][j] = true
				}
			}
		}
	}

	// 我们只想知道， dp[0][j] 的次数，其它的我们不关心， 哎！ 因此 dp 也可以降维了，dp[j] 就可以了

	// using dp to sovle the min cut
	for l := 0; l <= n; l++ {
		//for i := 0; i < n; i++ {
		i := 0
		j := i + l
		if j < n {
			if p[i][j] {
				dp[0][j] = 0
			} else {
				tmp := l + 1
				for k := i; k < j; k++ {
					if p[k+1][j] {
						tmp = min(tmp, dp[0][k]+1)
					}
				}
				dp[i][j] = tmp
			}
		}
		//}
	}

	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
