package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	n := len(nums)
	leftones := make([]int, n)
	rightones := make([]int, n)

	var getter = func(i int, counts []int) int {
		if i < 0 || i >= n {
			return 0
		}
		return counts[i]
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			leftones[i] = getter(i-1, leftones) + 1
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 1 {
			rightones[i] = getter(i+1, rightones) + 1
		}
	}

	ans := 0
	// special case all ones
	if leftones[n-1] == n {
		return n
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			ans = max(ans, getter(i-1, leftones)+getter(i+1, rightones)+1)
		} else {

		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
