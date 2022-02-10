package leetcode

import (
	"math"
)

/*
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/4-xun-zhao-liang-ge-zheng-xu-shu-zu-de-zhong-we-15/
这是我看到的最好的解法
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	left, right := 0, n
	var currA, currB int

	var getter = func(pos int, nums []int) int {
		if pos < 0 {
			return math.MinInt64
		}
		if pos >= len(nums) {
			return math.MaxInt64
		}
		return nums[pos]
	}
	for left <= right {

		currA = left + (right-left)>>1
		currB = (1+len(nums1)+len(nums2))/2 - currA

		//invariants is  LA < RB && LB < RA

		if getter(currA-1, nums1) > getter(currB, nums2) {
			right = currA - 1
		} else if getter(currB-1, nums2) > getter(currA, nums1) {
			left = currA + 1
		} else {
			// must found
			if (len(nums1)+len(nums2))%2 == 0 {
				ans := float64(max(getter(currA-1, nums1), getter(currB-1, nums2)))
				ans += float64(min(getter(currA, nums1), getter(currB, nums2)))
				return ans / 2.0
			}
			// odd
			return float64(max(getter(currA-1, nums1), getter(currB-1, nums2)))

		}
	}
	return -1
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
