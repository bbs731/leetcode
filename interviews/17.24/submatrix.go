package leetcode

import "math"

/*
这道题，没有看解题思路的时候，是做不出来的！
 */
func getMaxMatrix(matrix [][]int) []int {

	m := len(matrix)
	n := len(matrix[0])
	ans := math.MinInt64
	record := []int{}

	// 计算列的前缀和。
	sum := make([][]int, n)
	for j := 0; j < n; j++ {
		sum[j] = make([]int, m)
		sum[j][0] = matrix[0][j]
		for i := 1; i < m; i++ {
			sum[j][i] = sum[j][i-1] + matrix[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			b := make([]int, n)
			for k := 0; k < n; k++ {
				b[k] = sum[k][j] - sum[k][i] + matrix[i][k]
			}

			// 找 b[k] 序列的， 最大连续子数列的和
			begincol := 0
			if b[0] > ans {
				ans = b[0]
				record = []int{i, 0, j, 0}
			}
			for q := 1; q < n; q++ {
				if b[q-1]+b[q] > b[q] {
					b[q] += b[q-1]
				} else {
					begincol = q
				}
				if b[q] > ans {
					ans = b[q]
					record = []int{i, begincol, j, q}
				}
			}
		}
	}
	return record
}
