package leetcode

import "sort"

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	dp := make([]int, n)
	prev := make([]int, n)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		dp[i] = 1
		prev[i] = -1
	}
	ans := 1
	index := 0  // 初始化成 -1 和 0 是有区别的。

	for i := 1; i < n; i++ {
		c := 1 // 初始化， 1 和 0 是有区别的。
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				//c = max(c, dp[j]+1)
				if dp[j]+1 > c {
					c = dp[j] + 1
					prev[i] = j
				}
			}
		}
		dp[i] = c
		//ans = max(ans, c)
		if c > ans {
			ans = c
			index = i
		}
	}
	sets := make([]int, 0)
	for index >= 0 {
		sets = append(sets, nums[index])
		index = prev[index]
	}

	return sets
}
