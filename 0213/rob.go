package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	// do not select element 0 no head
	nh := make([]int, n)
	nh[0] = 0
	nh[1] = nums[1]
	nh[2] = max(nums[1], nums[2])

	// select element 0
	wh := make([]int, n)
	wh[0] = nums[0]
	wh[1] = nums[0]

	for i := 2; i < n; i++ {
		nh[i] = max(nh[i-2]+nums[i], nh[i-1])
		wh[i] = max(wh[i-2]+nums[i], wh[i-1])
	}

	return max(nh[n-1], wh[n-2])
}
