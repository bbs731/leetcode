package leetcode

func minimumReplacement(nums []int) int64 {

	n := len(nums)
	lastitem := nums[n-1]

	splits := 0

	for i := n - 2; i >= 0; i-- {

		if nums[i] <= lastitem {
			lastitem = nums[i]
			continue
		}

		d := nums[i] / lastitem

		// nums[i] > lastitem
		if nums[i] == d*lastitem {
			splits += d - 1
			// lastitem no change
		} else {
			start := 2
			end := d + 1

			// 二分查找， 永远是你的要害啊！
			for end > start {
				mid := (start + end) / 2
				if nums[i]/mid >= lastitem {
					start = mid + 1
				} else {
					end = mid
				}
			}

			splits += end - 1
			lastitem = nums[i] / end
		}
	}
	return int64(splits)
}
