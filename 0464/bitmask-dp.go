package leetcode

/*
这道题，太变态了， 完全想不到解法！

下次，重做 ！
 */
func canIWin(maxChoosableInteger int, desiredTotal int) bool {

	if maxChoosableInteger*(maxChoosableInteger+1) < 2*desiredTotal {
		return false
	}

	size := 1 << uint(maxChoosableInteger)
	// 1 stands for black win , -1 white win, 0 dual (not decided yet)
	dp := make([]int, size)

	var dfs func(state, total int) int

	dfs = func(state, total int) int {
		if dp[state] != 0 {
			return dp[state]
		}
		for k := 0; k < maxChoosableInteger; k++ {
			if state&(1<<uint(k)) != 0 {
				continue
			}

			if k+1 >= total || dfs(state|(1<<uint(k)), total-k-1) == 1 { // 我们成功了，或者，对手失败了！
				dp[state] = 1
				return 1
			}
		}
		dp[state] = -1
		return -1
	}

	if dfs(0, desiredTotal) == 1 {
		return true
	}
	return false
}
