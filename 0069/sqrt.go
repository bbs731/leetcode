package leetcode

func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)>>1
		n := mid * mid
		if n == x {
			return mid
		}
		if n > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
