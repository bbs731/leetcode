package leetcode

import "strconv"

func atMostNGivenDigitSet(digits []string, n int) int {

	nums := make(map[int]struct{})
	for _, d := range digits {
		nd, _ := strconv.Atoi(d)
		nums[nd] = struct{}{}
	}
	nums[0] = struct{}{} // adding digit 0

	l := 0
	digitstr := make([]string, 0)
	for n > 0 {
		m := n % 10
		digitstr = append(digitstr, strconv.Itoa(m))
		n = n / 10
		l++
	}

	dp := make([][]int, l)

	for i := 0; i < l; i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = -1
		dp[i][1] = -1
	}

	var getdp func(int, int, bool, []string, map[int]struct{}, [][]int) int

	/* 带前导零的 digit DP 模板 */
	getdp = func(pos int, limit int, lead bool, digits []string, sets map[int]struct{}, dp [][]int) int {
		if pos == -1 {
			return 1
		}
		if !lead && dp[pos][limit] != -1 {
			return dp[pos][limit]
		}
		var up int
		if limit == 1 {
			up, _ = strconv.Atoi(digits[pos])
		} else {
			up = 9
		}

		ans := 0
		for i := 0; i <= up; i++ {
			if _, ok := sets[i]; !ok {
				continue
			}
			if i == 0 && (!lead || pos == 0 && lead) {
				continue
			}

			islimit := 0
			if limit > 0 && i == up {
				islimit = 1
			}
			ans += getdp(pos-1, islimit, lead && i == 0, digits, sets, dp)
		}
		if !lead {
			dp[pos][limit] = ans
		}
		return ans
	}

	return getdp(l-1, 1, true, digitstr, nums, dp)
}
