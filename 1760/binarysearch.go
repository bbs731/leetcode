package leetcode

/*
type: binary search
 */
func minimumSize(nums []int, maxOperations int) int {

	largest := nums[0]
	for i := 0; i < len(nums); i++ {
		largest = max(largest, nums[i])
	}
	l := 1
	r := largest

	for l < r {
		mid := l + (r-l)>>1
		total := 0
		for i := 0; i < len(nums); i++ {
			total += floor(nums[i], mid)
		}

		if total > maxOperations {
			// should increase mid so
			l = mid + 1
		} else {
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

func floor(a, b int) int {
	// b must no be zero
	if a <= b {
		return 0
	}

	if a%b == 0 {
		return a/b - 1
	}
	return a / b
}
