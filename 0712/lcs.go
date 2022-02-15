package leetcode

/*
看看 1143 LCS （longest common sequence) 的解题，框架。

和这道题是一致的，不同的是 状态转移方程式
 */
func minimumDeleteSum(s1 string, s2 string) int {
	N1 := len(s1)
	N2 := len(s2)

	f := make([][]int, N1+1)
	for i := 0; i <= N1; i++ {
		f[i] = make([]int, N2+1)
	}

	// 初始化
	for i := 1; i <= N2; i++ {
		f[0][i] = f[0][i-1] + int(s2[i-1])
	}
	for i := 1; i <= N1; i++ {
		f[i][0] = f[i-1][0] + int(s1[i-1])
	}

	for i := 1; i <= N1; i++ {
		for j := 1; j <= N2; j++ {
			if s1[i-1] == s2[j-1] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = min(f[i-1][j]+int(s1[i-1]), f[i][j-1]+int(s2[j-1]))
			}
		}
	}

	//return f[N1-1][N2-1]
	return f[N1][N2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
