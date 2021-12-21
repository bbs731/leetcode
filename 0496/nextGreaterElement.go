package leetcode

/*
nums2 = [ 2,3,8,4, 5, 7, 1, 6, 9]
next  = [                     -1]
        [                   9,-1]
		[                6, 9,-1]
		[             9, 6, 9,-1]
		[          7, 9, 6, 9,-1]
		[       5, 7, 9, 6, 9,-1]
		[     9,5, 7, 9, 6, 9,-1]
		[   8,9,5, 7, 9, 6, 9,-1]
		[ 3,8,9,5, 7, 9, 6, 9,-1]


next 变现为一个单调栈。具体看官网： https://leetcode-cn.com/problems/next-greater-element-i/solution/xia-yi-ge-geng-da-yuan-su-i-by-leetcode-bfcoj/
 */

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	next := make([]int, len(nums2))
	m := make(map[int]int, len(nums2))

	next[len(nums2)-1] = -1
	m[nums2[len(nums2)-1]] = len(nums2) - 1
	for i := len(nums2) - 2; i >= 0; i-- {
		if nums2[i] < nums2[i+1] {
			next[i] = nums2[i+1]
		} else { // nums2[i] > nums2[i+1], given there is no equal numbers
			j := i + 1
			for ; nums2[i] > next[j] && next[j] != -1; j++ {
			}
			next[i] = next[j]
		}

		m[nums2[i]] = i
	}

	for i, v := range nums1 {
		pos := m[v]
		ans[i] = next[pos]
	}
	return ans
}
