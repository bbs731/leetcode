package leetcode

/*
我对二分查找的疑惑，
1. left <= right 和  left < right 如何选择？
2. mid = left + (right - left) >> 1 和  mid = (left + right +1) >> 1 这个如何选择呢?
 */
func arrangeCoins(n int) int {
	left, right := 0, n

	for left <= right {
		mid := left + (right-left)>>1
		total := (1 + mid) * mid >> 1

		if total == n {
			return mid
		}

		if total > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 这个 left - 1 是不是显得不优雅啊？
	return left - 1
}
