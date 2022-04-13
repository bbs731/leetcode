package leetcode

/*
这就是所谓的， 尺取法吗？
 */
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)

	prefixsum := make([]int, n)
	prefixsum[0] = nums[0]
	if prefixsum[0] >= target {
		return 1
	}
	var i int
	var right int
	right = -1
	for i = 1; i < n; i++ {
		prefixsum[i] = prefixsum[i-1] + nums[i]
		if right == -1 && prefixsum[i] >= target {
			right = i
		}
	}
	if right == -1 {
		return 0
	}

	ans := right - 0 + 1
	left := 0

	if right != n-1 {
		right++
	}

	for ; right < n; right++ {
		for left <= right && prefixsum[right]-prefixsum[left]+nums[left] >= target {
			ans = min(ans, right-left+1)
			left++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
