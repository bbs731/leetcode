package leetcode

import (
	"math"
	"sort"
)

/*
	这道题的难点， 不在于 DP 部分，

	在于如何计算 dist[i][j] 部分， 这里的计算， 这里的实现复杂度是 O（n^4) 还是 O(n ^3)， 官网上给了 O(n^2) 的解法，

	dist[i][j] = dist[i+1][j-1] + houses[j] - houses[i]

本题最终的时间复杂度是  O(k * n^2)

 */
func cost(houses []int, position int) int {
	total := 0
	for i := 0; i < len(houses); i++ {
		if houses[i] >= position {
			total += houses[i] - position
		} else {
			total += position - houses[i]
		}
	}
	return total
}

// 给定一个数组， 找到一个邮箱最优posistion, 使得 total distance 最小， 并返回 distance。
func findPosition(houses []int) int {
	n := len(houses)
	if n == 1 {
		return 0
	}
	min := math.MaxInt64
	for i := 0; i < len(houses); i++ {
		c := cost(houses, houses[i])
		if c < min {
			min = c
		}
	}
	return min
}

func minDistance(houses []int, k int) int {

	n := len(houses)
	if k >= n {
		return 0
	}
	sort.Slice(houses, func(i, j int) bool { return houses[i] < houses[j] })
	dp := make([][]int, n)
	// dist 保存的是 i, j 中间放一个 邮箱的 cost
	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][1] = findPosition(houses[:i+1])
		dist[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			dist[j][i] = findPosition(houses[j : i+1])
		}
	}

	for j := 2; j <= k; j++ {
		for i := 0; i < n; i++ {
			dp[i][j] = dp[i][j-1]
			for s := i - 1; s >= 0; s-- {
				//dp[i][j] = min(dp[i][j-1], dp[i-1][j-1]+rdp[i])
				dp[i][j] = min(dp[i][j], dp[s][j-1]+dist[s+1][i])
			}
		}
	}

	return dp[n-1][k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
