package leetcode

/*
DP 中最经典的例题之一了。 熟记吧！

算法导论里，有详细的解读 ！

 */
func longestCommonSubsequence(text1 string, text2 string) int {
	N1 := len(text1)
	N2 := len(text2)

	f := make([][]int, N1)
	for i := 0; i < N1; i++ {
		f[i] = make([]int, N2)
	}

	getter := func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		return f[i][j]
	}

	for i := 0; i < N1; i++ {
		for j := 0; j < N2; j++ {
			if text1[i] == text2[j] {
				f[i][j] = getter(i-1, j-1) + 1
			} else {
				f[i][j] = max(getter(i-1, j), getter(i, j-1))
			}
		}
	}

	return f[N1-1][N2-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
