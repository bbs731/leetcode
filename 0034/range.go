package leetcode

/*
试着， 学习一下，

sort.SearchInts 的用法！！！！

不过，就本题的解法来说， 做的完美， 二分查找，需要收敛的掌握，尤其，怎么处理边界条件 
 */
func searchRange(nums []int, target int) []int {

	// lower index
	lo, hi := 0, len(nums)-1
	low := -1
	for lo <= hi {
		mid := lo + (hi-lo+1)>>1
		if nums[mid] == target {
			low = mid
			hi = mid - 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	if low == -1 {
		return []int{-1, -1}
	}

	// upper index
	lo, hi = 0, len(nums)-1
	upper := low
	for lo <= hi {
		mid := lo + (hi-lo+1)>>1
		if nums[mid] == target {
			upper = mid
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return []int{low, upper}
}
