package leetcode

import "math"

/*
你也不是被，吓大的，咋这么怂呢？

这道题， 可以转化成一个 0-1 背包的 DP 问题！  用 DP的方法再解一遍吧！
 */
func lastStoneWeightII(stones []int) int {

	if len(stones) == 1 {
		return stones[0]
	}

	sum := make(map[int]struct{})
	sum[stones[0]] = struct{}{}
	sum[-stones[0]] = struct{}{}

	for i := 1; i < len(stones); i++ {
		tmp := make(map[int]struct{})
		for k := range sum {
			tmp[k+stones[i]] = struct{}{}
			tmp[k-stones[i]] = struct{}{}
		}
		sum = tmp
	}

	ans := math.MaxInt64 >> 1
	for k := range sum {
		if k >= 0 {
			ans = min(ans, k)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
