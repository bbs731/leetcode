package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxLen(nums []int) int {
	var ans, positiveL, negativeL int

	if nums[0] > 0 {
		positiveL = 1
	} else if nums[0] < 0 {
		negativeL = 1
	}

	ans = max(ans, positiveL)

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			positiveL++
			if negativeL > 0 {
				negativeL++
			}

		} else if nums[i] < 0 {
			prevn, prevp := negativeL, positiveL

			if prevn > 0 {
				positiveL = prevn + 1
			} else {
				positiveL = 0
			}

			if prevp > 0 {
				negativeL = prevp + 1
			} else {
				negativeL = 1
			}

		} else {
			positiveL = 0
			negativeL = 0
		}

		ans = max(ans, positiveL)
	}
	return ans
}
