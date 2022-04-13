package leetcode

import (
	"sort"
)

func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	counts := 0

	for a := 0; a < n-2; a++ {
		for b := a + 1; b < n-1; b++ {
			c := nums[a] + nums[b]
			pos := sort.SearchInts(nums[b+1:n], c)
			//pos := sort.SearchInts(nums, c)
			//if pos == n {
			//	pos = n - 1
			//}
			//if nums[pos] >= c {
			//	pos--
			//}
			//if pos > b {
			//	counts += pos - b
			//}
			counts += pos
		}
	}
	return counts
}
