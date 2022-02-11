package leetcode

import "sort"
/*
这道题，对目前我还是太难了！
 */
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	// 在 0 和 max distance 之间找一个数 mid ，然后统计 nums[i], nums[j] 有多少个满足  nums[j] - nums[i] <=  mid
	i, j := 0, nums[len(nums)-1]-nums[0]
	for i < j {
		mid := i + (j-i)>>1 // this is a guess number

		left := 0
		count := 0
		// 这是一个双指针的问题， 在 O(n) 的时间内，做了统计
		for right := 0; right < len(nums); right++ {
			for nums[right]-nums[left] > mid {
				left++
			}
			count += right - left
		}

		if count >= k {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}
