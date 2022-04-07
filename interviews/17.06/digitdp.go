package leetcode

/* 和官网的 233 一个解法

时间复杂度是 Log(n)
 */
func numberOf2sInRange(n int) int {

	ans := 0
	for k, mulk := 0, 1; n >= mulk; k++ {
		ans += (n/(mulk*10))*mulk + min(max(n%(mulk*10)-2*mulk+1, 0), mulk)
		mulk *= 10
	}
	return ans
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
