package leetcode

/*
两种 dfs 的方法， 一种是 mask hat , 复杂度是 10* 2^40
一种是 mask persion, 复杂度是 40 * 2^10,
无论那种方法，都需要 memory based dfs,  但是 明显  2^40 的方法， memory 使用会超过限制！
 */
func numberWays(hats [][]int) int {
	n := len(hats)
	mod := 1000000000 + 7

	maxid := 0
	for i := 0; i < n; i++ {
		for _, h := range hats[i] {
			maxid = max(maxid, h)
		}
	}

	hatstoperson := make([][]int, maxid+1)

	for i := 0; i < n; i++ {
		for _, h := range hats[i] {
			hatstoperson[h] = append(hatstoperson[h], i)
		}
	}

	target := 1<<uint(n) - 1
	dp := make([][]int, target+1)
	for i := 0; i < target+1; i++ {
		dp[i] = make([]int, maxid+1)
		for j := 0; j < maxid+1; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(int, uint, uint) int
	// bitmask  用来记录人  n <=10
	dfs = func(hat int, bitmask uint, target uint) int {
		if bitmask == target {
			return 1
		}
		if hat == maxid+1 {
			return 0
		}
		if dp[bitmask][hat] != - 1 {
			return dp[bitmask][hat]
		}
		total := dfs(hat+1, bitmask, target)
		for _, p := range hatstoperson[hat] {
			bit := uint(1 << uint(p))
			if bit&bitmask == 0 {
				total += dfs(hat+1, bit|bitmask, target)
				total %= mod
			}
		}
		dp[bitmask][hat] = total
		return total
	}
	return dfs(1, 0, 1<<uint(n)-1) % mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
