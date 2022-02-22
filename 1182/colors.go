package leetcode
/*
这个代码好臃肿啊！
 */
func shortestDistanceColor(colors []int, queries [][]int) []int {
	dp := make([][]int, len(colors))

	yellow := -1 // 1
	blue := -1   //2
	red := -1    //3
	for i := 0; i < len(colors); i++ {
		dp[i] = make([]int, 4)
		for j := 1; j <= 3; j++ {
			dp[i][j] = -1
		}
	}

	for i := 0; i < len(colors); i++ {
		if colors[i] == 1 {
			yellow = i
		} else if colors[i] == 2 {
			blue = i
		} else if colors[i] == 3 {
			red = i
		}

		if yellow != -1 {
			dp[i][1] = i - yellow
		}
		if blue != -1 {
			dp[i][2] = i - blue
		}
		if red != -1 {
			dp[i][3] = i - red
		}
	}

	yellow = -1 // 1
	blue = -1   //2
	red = -1    //3

	for i := len(colors) - 1; i >= 0; i-- {
		if colors[i] == 1 {
			yellow = i
		} else if colors[i] == 2 {
			blue = i
		} else {
			red = i
		}

		if yellow != -1 {
			dp[i][1] = min(dp[i][1], yellow-i)
		}
		if blue != -1 {
			dp[i][2] = min(dp[i][2], blue-i)
		}
		if red != -1 {
			dp[i][3] = min(dp[i][3], red-i)
		}
	}

	ans := make([]int, 0)
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		ans = append(ans, dp[query[0]][query[1]])
	}
	return ans
}

func min(a, b int) int {
	if a == -1 {
		return b
	}
	if a < b {
		return a
	}
	return b
}
