package leetcode

import "sort"

// 这题， 不看题解， 做不出来啊！ 

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var p int
	results := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			p = len(nums) - 1
			for j := i + 1; j < len(nums); j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					p = len(nums) -1
					for k := j + 1; k < len(nums); k++ {
						if k == j+1 || nums[k] != nums[k-1] {
							for k < p && nums[i]+nums[j]+nums[k]+nums[p] > target {
								p = p - 1
							}
							if k == p {
								break
							}
							if nums[i]+nums[j]+nums[k]+nums[p] == target {
								results = append(results, []int{nums[i], nums[j], nums[k], nums[p]})

							}
						}
					}
				}
			}
		}
	}
	return results
}
