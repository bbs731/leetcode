package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {

	n := len(nums)
	// 不是最优解， space 可以省， 可以只用一维的数组表示
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
		c[i][i] = nums[i]
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			s1, s2 := -1, -1

			if i+2 <= j {
				s1 = c[i][i] + c[i+2][j]
			} else {
				s1 = c[i][i]
			}
			s2 = c[i+1][j]

			c[i][j] = max(s1, s2)
		}
	}

	return c[0][n-1]
}
