package leetcode

/*
 partition 快速排序中的 partition 方法类似
 */
func findKthLargest(nums []int, k int) int {

	pos := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[0] {
			continue
		}
		nums[pos+1], nums[i] = nums[i], nums[pos+1]
		pos++
	}
	nums[pos], nums[0] = nums[0], nums[pos]

	if len(nums)-pos == k {
		return nums[pos]
	}

	if len(nums)-pos > k {
		return findKthLargest(nums[pos+1:], k)
	} else {
		return findKthLargest(nums[:pos], k-len(nums)+pos)
	}
}
