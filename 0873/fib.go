package leetcode

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	ans := 0
	index := make(map[int]int)
	for i := 0; i < n; i++ {
		index[arr[i]] = i
	}

	for i := 2; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			numk := arr[i] - arr[j]
			if k, ok := index[numk]; ok && k < j {
				dp[i][j] = dp[j][k] + 1
				ans = max(ans, dp[i][j])
			}
		}
	}

	if ans > 0 {
		return ans + 2
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
