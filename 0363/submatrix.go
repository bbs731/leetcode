package leetcode

import "math"
/*
wow, 有了 《程序员面试宝典6》 17.24 题目的经验， 这个就做的够快了！

 */
func maxSumSubmatrix(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])

	// 计算列的前缀和
	colsum := make([][]int, n)
	for j := 0; j < n; j++ {
		colsum[j] = make([]int, m)
		colsum[j][0] = matrix[0][j]
		for i := 1; i < m; i++ {
			colsum[j][i] = colsum[j][i-1] + matrix[i][j]
		}
	}

	b := make([]int, n)
	bsum := make([]int, n)
	ans := math.MaxInt64

	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			for k := 0; k < n; k++ {
				b[k] = colsum[k][j] - colsum[k][i] + matrix[i][k]
			}
			// now 我们来计算 b[k] 序列的前缀和
			bsum[0] = b[0]
			for q := 1; q < n; q++ {
				bsum[q] = bsum[q-1] + b[q]
			}
			near := nearest(bsum, k)
			ans = min(ans, near)
		}
	}
	return k - ans
}

func nearest(bsum []int, k int) int {
	near := math.MaxInt64
	n := len(bsum)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			var sum int
			if i-1 >= 0 {
				sum = bsum[j] - bsum[i-1]
			} else {
				sum = bsum[j]
			}
			if sum <= k {
				near = min(near, k-sum)
			}
		}
	}
	return near
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
