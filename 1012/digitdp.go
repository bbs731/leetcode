package leetcode

/*
经典的 digit DP 问题。 记住这个模板！

1. 有前导零
2. 有state (本例子是 isOk, 代表是否出现了重复数字）
3. dp cache 的问题， 这里面需要换成， pos, state, visited 的三元组，  在不考虑 Lead 和 limit 这种只出现一次的情况！
 */
func numDupDigitsAtMostN(n int) int {
	m := n
	digits := make([]int, 0)
	for m > 0 {
		digits = append(digits, m%10)
		m = m / 10
	}

	dp := make([][][]int, len(digits))

	size := 1024 // 因为 n < 10^9     1<<9 = 512
	for i := 0; i < len(digits); i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j <= 1; j++ {
			dp[i][j] = make([]int, size)
			for k := 0; k < size; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	var getdp func(int, int, bool, int, uint, [][][]int) int

	/* 不带前导零的 digit DP 模板 */
	getdp = func(pos int, limit int, lead bool, isOk int, visited uint, dp [][][]int) int {
		if pos == -1 {
			if isOk > 0 {
				return 1
			} else {
				return 0
			}
		}
		if !lead && limit == 0 && dp[pos][isOk][visited] != -1 {
			return dp[pos][isOk][visited]
		}
		var up int
		if limit == 1 {
			up = digits[pos]
		} else {
			up = 9
		}

		ans := 0
		for i := 0; i <= up; i++ {
			islimit := 0
			if limit > 0 && i == up {
				islimit = 1
			}

			if lead && i == 0 {
				ans += getdp(pos-1, islimit, true, isOk, visited, dp)

			} else {
				shouldOk := isOk
				has := visited & (1 << uint(i))
				if has != 0 {
					shouldOk = 1
				}
				ans += getdp(pos-1, islimit, false, shouldOk, visited|(1<<uint(i)), dp)
			}
		}
		if !lead && limit == 0 {
			dp[pos][isOk][visited] = ans
		}
		return ans
	}

	return getdp(len(digits)-1, 1, true, 0, 0, dp)

}
