package leetcode

/*
时间复杂度，怎么求？
O（n^2 *k)  怎么算的？
 */
func maxVacationDays(flights [][]int, days [][]int) int {

	n := len(days)
	k := len(days[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k)
		for j := 0; j < k; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(pos int, week int) int {
		if week == k {
			return 0
		}
		if dp[pos][week] != -1 {
			return dp[pos][week]
		}
		// spend current week in pos city
		accum := days[pos][week]
		rest := 0
		// for next week
		for i := 0; i < n; i++ {
			if flights[pos][i] == 1 {
				rest = max(rest, dfs(i, week+1))
			}
		}
		// no place to go, stay at city pos is an alternative as well
		rest = max(rest, dfs(pos, week+1))

		dp[pos][week] = accum + rest
		return dp[pos][week]
	}

	ans := 0
	for i := 0; i < n; i++ {
		if flights[0][i] == 1 {
			ans = max(ans, dfs(i, 0))
		}
	}
	ans = max(ans, dfs(0, 0))
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
