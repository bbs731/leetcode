package leetcode

/*
 行吧， 这两道题的解法，直接，证明了你的智商不行！
 */
func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		// 把所有的 invalid number 变成  n+1 是一个正数，（ > 0 很关键）
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		x := nums[i]
		if x < 0 {
			x = -x
		}
		// 如果 nums[i] 是 【1,n ] 之间的数， 把它对应的在index 的数，在数组中 Mark 成负的。
		if x >= 1 && x <= n {
			if nums[x-1] > 0 {
				nums[x-1] = -nums[x-1]
			}
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i+1
		}
	}

	return n + 1
}
