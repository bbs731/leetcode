package leetcode

/*
我觉得这就是一个简单的 dfs,  熟练 mask 在这里面的用法！ 这是最基本的 bitmask DP 的形态了！
没有时间复杂度的要求， n <=8  不涉及到优化！
 */

func numberOfPatterns(m int, n int) int {

	cross := make([][]int, 10)
	for i := 0; i < 10; i++ {
		cross[i] = make([]int, 10)
	}
	// 初始化
	cross[1][9] = 5
	cross[9][1] = 5
	cross[1][3] = 2
	cross[3][1] = 2
	cross[1][7] = 4
	cross[7][1] = 4
	cross[2][8] = 5
	cross[8][2] = 5
	cross[3][9] = 6
	cross[9][3] = 6
	cross[3][7] = 5
	cross[7][3] = 5
	cross[4][6] = 5
	cross[6][4] = 5
	cross[7][9] = 8
	cross[9][7] = 8

	var dfs func(uint, int, int) int

	dfs = func(start uint, n int, mask int) int {
		if n == 0 {
			return 1
		}
		ans := 0

		var i uint
		for i = 1; i <= 9; i++ {
			if i == start {
				continue
			}

			if (mask & (1 << i)) != 0 { // visited i before
				continue
			}
			cp := uint(cross[start][i])
			if cp == 0 || (mask&(1<<cp)) != 0 { // no cross point or cross point is visited before
				ans += dfs(i, n-1, mask|(1<<i))
			}
		}
		return ans
	}
	cnts := 0
	var k uint
	for k = 1; k <= 9; k++ {
		for j := m; j <= n; j++ {
			cnts += dfs(k, j-1, 1<<k)
		}
	}
	return cnts
}
