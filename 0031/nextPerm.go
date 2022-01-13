package leetcode

import "sort"

/*
 	第一遍思路还是有漏洞啊。  // [2, 3, 1] 就没考虑清楚
 */
func nextPermutation(nums []int) {
	N := len(nums)
	j := N - 1
	for ; j > 0; j-- {
		if nums[j-1] < nums[j] {
			break
		}
	}

	if j == 0 {
		// speical case, need to reverse the nums would be the answer
		// 但是我们和  else 大算用一样的处理方法
		nums[0], nums[N-1] = nums[N-1], nums[0]

	} else {
		// [2, 3, 1]
		k := N - 1
		for ; k >= j; k-- {
			if nums[k] > nums[j-1] {
				break
			}
		}
		nums[j-1], nums[k] = nums[k], nums[j-1]

	}
	sort.Ints(nums[j:N])
}
