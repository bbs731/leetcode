package leetcode

func smallestDivisor(nums []int, threshold int) int {
	sum := 0
	largest := nums[0]

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		largest = max(largest, nums[i])
	}

	l := 1
	r := largest

	for l < r {
		mid := l + (r-l)>>1

		total := 0
		for i := 0; i < len(nums); i++ {
			total += max(ceiling(nums[i], mid), 1)
		}

		if total > threshold {
			l = mid + 1
		} else if total <= threshold {

			r = mid
		}
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ceiling(a, b int) int {
	if b == 0 {
		return 0
	}
	if a%b == 0 {
		return a / b
	}
	return a/b + 1
}
