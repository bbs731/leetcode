package leetcode

import "fmt"

type rect struct {
	w int
	h int
}

/*
这个算法是有漏洞的， 找别的办法把！
 */
func maximalRectangle(matrix [][]byte) int {

	m := len(matrix)    // number of rows
	n := len(matrix[0]) // cols

	ans := 0

	dp := make([][][]rect, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]rect, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]rect, 2)
		}
	}

	// 初始化 first w and first h
	if matrix[0][0] == '1' {
		dp[0][0][0] = rect{1, 1} // 水平反向的矩形
		dp[0][0][1] = rect{1, 1} // 垂直方向
		ans = 1
	}

	// first row
	for i := 1; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i][0] = rect{dp[0][i-1][0].w + 1, 1}
			dp[0][i][1] = rect{dp[0][i][0].w, 1}
			ans = max(ans, dp[0][i][0].w)
		} // else 都是 0 已经初始化了
	}
	// first col
	for i := 1; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0][1] = rect{1, dp[i-1][0][1].h + 1}
			dp[i][0][0] = rect{1, dp[i][0][1].h}
			ans = max(ans, dp[i][0][1].h)
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				//dp[i][j][0] 由 dp[i-1][j-1][0], dp[i-1][j][0] 和 dp[i][j-1][0] 的三个矩形重叠的地方决定

				if matrix[i-1][j-1] == '0' {
					dp[i][j][0] = rect{dp[i][j-1][0].w + 1, 1}
					dp[i][j][1] = rect{1, dp[i-1][j][1].h + 1}

				} else if matrix[i][j-1] == '0' {
					dp[i][j][1] = rect{1, dp[i-1][j][1].h + 1}
					dp[i][j][0] = rect{1, 1}

				} else if matrix[i-1][j] == '0' {
					dp[i][j][0] = rect{dp[i][j-1][0].w + 1, 1}
					dp[i][j][1] = rect{1, 1}

				} else {

					// 先计算 dp[i][j-1][0] 和  dp[i-1][j-1][0] 连个矩形相交的矩形的长和宽, 以 i，j-1 为右下角
					w := min(dp[i][j-1][0].w, dp[i-1][j-1][0].w)
					h := max(dp[i][j-1][0].h, dp[i-1][j-1][0].h+1) // 因为以  (i, j-1 ) 为右下角，所以 dp[i-1][j-1][0].h 需要加 1

					// 然后和  （i-1,j ) 为右下角的 rect 做 merge .  以 （i， j) 为 merge 过之后 矩形的右下角
					w = w + 1
					//w = min(w+1, dp[i-1][j][0].w)
					h = min(h, dp[i-1][j][0].h+1)

					dp[i][j][0] = rect{w, h}

					// 计算 dp[i][j][1]
					// 先 merge dp[i-1][j] 和 dp[i-1][j-1] 去 w 最大 ， h 最小的。
					h = min(dp[i-1][j][1].h, dp[i-1][j-1][1].h)
					w = max(dp[i-1][j][1].w, dp[i-1][j-1][1].w+1)

					w = min(w, dp[i][j-1][1].w+1)
					//h = min(h+1, dp[i][j-1][1].h)
					h = h + 1
					dp[i][j][1] = rect{w, h}

					// speical case 考虑， 但 i row and  j col
					ans = max(ans, dp[i-1][j][1].h+1)
					ans = max(ans, dp[i][j-1][0].w+1)

				}

				ans = max(ans, dp[i][j][0].w*dp[i][j][0].h)
				ans = max(ans, dp[i][j][1].w*dp[i][j][1].h)
			}
		}
	}
	print(m, n, dp, 0)
	return ans
}

func print(m, n int, dp [][][]rect, d int) {

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d x %d\t", dp[i][j][d].w, dp[i][j][d].h)
		}
		fmt.Println("")
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
