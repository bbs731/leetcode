package leetcode

/*
digit dp 模板和代码， 参见：
https://leetcode.cn/problems/digit-count-in-range/solution/leetcodeshu-wei-dpzong-jie-by-pedantic-einstein/

时间复杂度， 应该是 O(pos * pos * 10)   其中 pos = O(lg n)
 */
func countDigitOne(n int) (ans int) {
	x := n
	pos := 0
	a := make([]int, 20)
	dp := make([][]int, 20)
	for i := 0; i < 20; i++ {
		dp[i] = make([]int, 20)
		for j := 0; j < 20; j++ {
			dp[i][j] = -1
		}
	}
	for x != 0 {
		a[pos] = x % 10
		pos++
		x /= 10
	}

	var dfs func(int, bool, int) int
	dfs = func(pos int, limit bool, pre int) int {
		if pos == -1 {
			return pre
		}

		if !limit && dp[pos][pre] != -1 {
			return dp[pos][pre]
		}

		ans := 0
		up := 9
		if limit {
			up = a[pos]
		}

		for i := 0; i <= up; i++ {
			if i == 1 {
				ans += dfs(pos-1, limit && up == i, pre+1)
			} else {
				ans += dfs(pos-1, limit && up == i, pre)
			}
		}
		if !limit {
			dp[pos][pre] = ans
		}
		return ans
	}

	return dfs(pos-1, true, 0)
}
