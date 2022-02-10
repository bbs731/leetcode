package leetcode

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		h := left + (right-left)>>1 // prevent left+right overflow
		if nums[h] == target {
			return h
		} else if nums[h] > target {
			right = h - 1
		} else {
			left = h + 1
		}
	}
	// end condition left > right
	return -1
}
