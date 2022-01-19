package leetcode

func increasingTriplet(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	keep := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if keep[len(keep)-1] < nums[i] {
			keep = append(keep, nums[i])
			if len(keep) >= 3 {
				// early return
				return true
			}
		} else {
			// find the first element which greater or equal to  nums[i] and replace that element with nums[i]
			lo := 0
			hi := len(keep) - 1
			pos := 0
			for lo <= hi {
				mid := lo + (hi-lo+1)>>1
				if keep[mid] == nums[i] {
					pos = mid
					break
				} else if keep[mid] > nums[i] {
					pos = mid
					hi = mid - 1
				} else {
					lo = mid + 1
				}
			}
			keep[pos] = nums[i]
		}
	}

	return len(keep) >= 3
}
