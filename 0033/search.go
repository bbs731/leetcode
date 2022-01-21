package leetcode


/*
这里的逻辑， 我感觉能把人给绕死。

binary search 要精通， 多看这道题的解法！

 */
func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left+1)>>1
		if nums[mid] == target {
			return mid
		}

		// 这里面的套路好深啊！
		if target > nums[mid] {
			if nums[mid] > nums[left] {
				left = mid + 1
			} else {
				if target == nums[left] {
					return left
				}
				if target > nums[left] {
					right = mid - 1
				} else {
					left = mid + 1
				}

			}

		} else {
			// now target < nums[mid]
			if nums[mid] > nums[left] {
				if target == nums[left] {
					return left
				}
				if target > nums[left] {
					right = mid - 1
				} else {
					left = mid + 1
				}

			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
