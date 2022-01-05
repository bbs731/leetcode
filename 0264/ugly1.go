package leetcode

/*
经典题目，
没一次能答对
 */
func nthUglyNumber(n int) int {

	dp := make([]int, n+1)

	dp[0] = 1
	p2, p3, p5 := 1, 1, 1

	for i := 1; i <= n; i++ {

		c2, c3, c5 := p2*2, p3*3, p5*5
		c := min(min(c2, c3), c5)
		dp[i] = c

		if c == c2 {
			p2++
		}
		if c == c3 {
			p3++
		}
		if c == c5 {
			p5++
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
