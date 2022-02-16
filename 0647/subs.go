package leetcode

func countSubstrings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	var getter = func(i, j int) bool {
		// 比较 boundary 是错误的
		//if i < 0 || i >= n {
		//	return true
		//}
		//if j < 0 || j >= n {
		//	return true
		//}
		if i >= j {
			return true
		}
		return dp[i][j]
	}
	for l := 1; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			j := l + i - 1
			if s[i] == s[j] {
				dp[i][j] = getter(i+1, j-1)
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {

			if dp[i][j] {
				ans++
			}
		}
	}
	return ans
}
