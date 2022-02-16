package leetcode

import "sort"

// 这是一道 回溯的题目， 不是 DP
func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	sort.Ints(candidates)
	ans := make([][]int, 0)

	var dfs func(int, []int, int)
	dfs = func(k int, series []int, sum int) {
		if sum > target {
			return
		}
		if k >= n {
			return
		}
		if sum == target {
			// 这里需要 copy series, 然后存放在 ans 里， 不然 series 会变。
			l := make([]int, len(series))
			copy(l, series)
			ans = append(ans, l)
			return
		}
		if sum+candidates[k] > target {
			return
		}
		series = append(series, candidates[k])
		dfs(k, series, sum+candidates[k])
		// restore series
		series = series[:len(series)-1]

		dfs(k+1, series, sum)
	}

	dfs(0, []int{}, 0)
	return ans
}
