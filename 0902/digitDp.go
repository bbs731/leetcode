package leetcode

import (
	"strconv"
)

func atMostNGivenDigitSet(digits []string, n int) int {

	nums := make(map[int]struct{})
	for _, d := range digits {
		nd, _ := strconv.Atoi(d)
		nums[nd] = struct{}{}
	}
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

	var getdp func(int, int, []string, map[int]struct{}, [][]int) int

	/* 不带前导零的 digit DP 模板 */
	getdp = func(pos int, limit int, digits []string, sets map[int]struct{}, dp [][]int) int {
		if pos == -1 {
			return 1
		}
		if dp[pos][limit] != -1 {
			return dp[pos][limit]
		}
		dp[pos][limit] = 0
		var up int
		if limit == 1 {
			up, _ = strconv.Atoi(digits[pos])
		} else {
			up = 9
		}

		for i := range sets {
			if i > up {
				continue
			}

			islimit := 0
			if limit > 0 && i == up {
				islimit = 1
			}
			dp[pos][limit] += getdp(pos-1, islimit, digits, sets, dp)
		}
		return dp[pos][limit]
	}

	/* 这里，是不是，太恶心了！ */

	ans := getdp(l-1, 1, digitstr, nums, dp)
	for k := l - 2; k > 0; k-- {
		ans += getdp(k, 0, digitstr, nums, dp)
	}
	return ans

}
