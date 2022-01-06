package leetcode

import "math"

func findPeakElement(nums []int) int {
	n := len(nums)
	low, high := 0, len(nums)-1

	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	for low <= high {
		mid := low + (high-low)>>1

		if get(mid) > get(mid+1) && get(mid) > get(mid-1) {
			return mid
		}
		if get(mid) > get(mid+1) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
