package leetcode

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	pp, p := 0, nums[0]

	for i := 1; i < len(nums); i++ {
		pp, p = p, max(pp+nums[i], p)
	}
	return p
}

func deleteAndEarn(nums []int) int {
	sort.Ints(nums)
	cost := make([]int, 0)

	acc := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			acc += nums[i]
		} else {
			cost = append(cost, acc)
			if nums[i] != nums[i-1]+1 {
				cost = append(cost, 0)
			}
			acc = nums[i]
		}
	}
	cost = append(cost, acc)

	return rob(cost)
}
