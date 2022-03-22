package leetcode

import (
	"fmt"
	"sort"
)

/*
	DP 状态压缩的典型，问题， 用DP 再做一遍！
 */
func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	sum := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] >= nums[j]
	})
	fmt.Println(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k

	if nums[0] > target {
		return false
	}

	sums := make([]int, k)
	used := make([]bool, n)

	var dfs func(int, int) bool
	dfs = func(bucket, level int) bool {
		if bucket == k {
			return true
		}
		if sums[bucket] == target {
			return dfs(bucket+1, 0)
		}

		for i := level; i < n; i++ {
			if used[i] {
				continue
			}
			if sums[bucket]+nums[i] > target {
				continue
			}
			used[i] = true
			sums[bucket] += nums[i]
			if dfs(bucket, level+1) {
				return true
			}
			// restore
			sums[bucket] -= nums[i]
			used[i] = false
			/*
				这个剪枝 有用啊， 要不然，时间会超时啊！
			 */
			for i < n && nums[i+1] == nums[i] {
				i++
			}
		}
		return false
	}
	return dfs(0, 0)
}
