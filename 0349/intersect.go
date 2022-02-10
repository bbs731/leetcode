package leetcode

import "sort"

func intersection(nums1 []int, nums2 []int) []int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	ans := make([]int, 0)
	var pos1, pos2 int
	for pos1 < len(nums1) && pos2 < len(nums2) {
		if nums1[pos1] == nums2[pos2] {
			// need to check duplicates
			if len(ans) == 0 || ans[len(ans)-1] != nums1[pos1] {
				ans = append(ans, nums1[pos1])
			}
			pos1++
			pos2++
		} else if nums1[pos1] < nums2[pos2] {
			pos1++
		} else {
			pos2++
		}
	}
	return ans
}
