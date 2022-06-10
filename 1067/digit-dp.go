package leetcode

/*
 dfs transition 计数的代码： 参考了这里：
https://leetcode.cn/problems/digit-count-in-range/solution/shu-wei-dp-by-cdyy_hai-2/


if lead && i ==0  {
	ans += dfs(pos-1, limit && up ==i, true, pre)
}
 */
func solve(n, d int) int {

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

	var dfs func(int, bool, bool, int) int
	dfs = func(pos int, limit bool, lead bool, pre int) int {
		if pos == -1 {
			return pre
		}

		if !limit && !lead && dp[pos][pre] != -1 {
			return dp[pos][pre]
		}

		ans := 0
		up := 9
		if limit {
			up = a[pos]
		}

		for i := 0; i <= up; i++ {
			if i == d {
				if lead && i == 0 {
					ans += dfs(pos-1, limit && up == i, true, pre)

				} else {
					ans += dfs(pos-1, limit && up == i, lead && i == 0, pre+1)
				}
			} else {
				ans += dfs(pos-1, limit && up == i, lead && i == 0, pre)
			}
		}
		if !limit && !lead {
			dp[pos][pre] = ans
		}
		return ans
	}

	return dfs(pos-1, true, true, 0)

}
func digitsCount(d int, low int, high int) int {
	return solve(high, d) - solve(low-1, d)
}
