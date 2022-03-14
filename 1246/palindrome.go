package leetcode

import (
	"fmt"
	"math"
)

func minimumMoves(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		// 初始化
		for j := 0; j < n; j++ {
			if i == j {
				dp[i][j] = 1
			} else if i > j {
				dp[i][j] = 0
			} else {
				dp[i][j] = math.MaxInt64 >> 1
			}
		}
	}
	fmt.Println(dp)
	// dp[i][j] = min { dp[i+1][j] + 1 , dp[i][j-1]+1  in case  arr[i] != arr[j] and  dp[i+1][j-1] in case arr[i] == arr[j] // and other k possible positions

	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			j := l - 1 + i
			if j < n {
				dp[i][j] = min(dp[i][j], min(dp[i+1][j], dp[i][j-1])+1) // 这里特殊处理了 l = 1 的情况
				if arr[i] == arr[j] {
					// 这里是特殊的处理。
					if dp[i+1][j-1] == 0 {
						dp[i][j] = 1
					} else {
						dp[i][j] = min(dp[i][j], dp[i+1][j-1])
					}
				}
				for k := i + 1; k <= j-1; k++ {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
				}
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
