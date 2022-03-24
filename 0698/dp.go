package leetcode

import (
	"sort"
)

/*
bitmask DP problem example


太牛了， 你学会了。

用 bitmask DP 的好处，就是， DP 是天然剪枝的,因此不需要什么技巧就能过时间的限制！ bitmask DP 也因此比朴素的搜索， 譬如 DFS search 要快
如果用 DFS， 你需要很强的剪枝技巧! 参看 kset.go 中 DFS 的解答！
 */

func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	sort.Ints(nums)

	total := 0
	for i := 0; i < n; i++ {
		total += nums[i]
	}

	// general check
	if total%k != 0 {
		return false
	}
	target := total / k
	// general check
	if nums[n-1] > target {
		return false
	}

	size := 1 << uint(n) // 最大是 2^16
	sums := make([]int, size)
	dp := make([]bool, size)
	dp[0] = true

	for i := 0; i < size; i++ {
		if dp[i] == false {
			continue
		}

		for j := 0; j < n; j++ {
			if i&(1<<uint(j)) != 0 {
				// 状态 i 已经有 nums[j] 了
				continue
			}

			if sums[i]%target+nums[j] <= target {
				dp[i|1<<uint(j)] = true
				sums[i|1<<uint(j)] = sums[i] + nums[j]
			} else {
				// since nums is sorted
				break
			}
		}
	}

	return dp[size-1]
}
