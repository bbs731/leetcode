package leetcode

func findKthLargest(nums []int, k int) int {

	left, right := 0, len(nums)-1
	target := len(nums) - k // the target position

	for left < right {
		mid := left + (right-left)>>1
		// use mid as pivot and do the partition, 之后再和 mid 没有关系了哈。 计算 pos
		nums[left], nums[mid] = nums[mid], nums[left]
		pos := left
		for i := left; i <= right; i++ {
			if nums[i] >= nums[left] {
				continue
			}
			pos++
			nums[pos], nums[i] = nums[i], nums[pos]
		}
		nums[pos], nums[left] = nums[left], nums[pos]
		if pos-left == target {
			return nums[pos]
		} else if pos-left > target {
			right = pos - 1
		} else {
			target -= pos - left + 1 // 这个顺序太微妙了， 先更新 target 再更新 left, target 用到了 left 做更新
			left = pos + 1
		}
	}
	return nums[left]
}
