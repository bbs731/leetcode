package leetcode

func isPerfectSquare(num int) bool {

	left, right := 0, num

	for left <= right {
		mid := left + (right-left)>>1
		n := mid * mid

		if n == num {
			return true
		}
		if n < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
