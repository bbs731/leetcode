package leetcode

/*
太难了， 做不出来！
 */
func maxCoins(nums []int) int {
	n := len(nums)

	vals := make([]int, n+2)

	vals[0] = 1
	vals[n+1] = 1
	for i := 1; i <= n; i++ {
		vals[i] = nums[i-1]
	}

	rect := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		rect[i] = make([]int, n+2)

		for j := 0; j < len(rect[i]); j++ {
			rect[i][j] = -1
		}
	}

	return solve(0, n+1, vals, rect)

}

func solve(left, right int, vals []int, rect [][]int) int {
	if left >= right-1 {
		return 0
	}
	if rect[left][right] != -1 {
		return rect[left][right]
	}
	for i := left + 1; i < right; i++ {
		sum := vals[left] * vals[i] * vals[right]
		sum += solve(left, i, vals, rect) + solve(i, right, vals, rect)
		rect[left][right] = max(rect[left][right], sum)
	}
	return rect[left][right]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
