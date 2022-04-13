package leetcode

/*
 这个方法，是错误的， 首先， 时间满足不了， 这是 O(n^2) 的解法。
其次， C[n][k] 很容易就 memoy overflow
 */
func findDerangement(n int) int {

	// dp[i] 代表有 i 个数，都不在自己的位置上的组合个数.  通过数学的归纳，知道 dp[i+1] 和 dp[i] 之间的关系
	dp := make([]int, max(4, n+1))
	dp[0] = 1
	dp[1] = 0
	dp[2] = 1 // 有两个数， 两个位置， 两个数都不在自己的位置的个数
	dp[3] = 2
	mod := 1000000007

	C := make([][]int, n+1) // C[n][k] = C[n-1][k-1] + C[n-1][k]
	for i := 0; i <= n; i++ {
		C[i] = make([]int, n+1)
	}
	C[0][0] = 1
	C[1][0] = 1
	C[1][1] = 1
	for i := 2; i <= n; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for k := 1; k <= i/2; k++ {
			C[i][k] = (C[i-1][k] + C[i-1][k-1]) % mod // 这里，在 i 足够大的时候，会 Overflow, 需要 mod
			C[i][i-k] = C[i][k]
		}
	}

	f := make([]int, max(4, n+1))
	f[0] = 1
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] * i
		f[i] %= mod
	}

	sol := make([]int, max(4, n+1))
	sol[0] = 1
	sol[1] = 0
	sol[2] = 1
	sol[3] = 2
	for i := 4; i <= n; i++ {
		ans := f[i]
		for k := i; k >= 0; k-- {
			ans = ans - C[i][i-k]*sol[k]
			ans = (ans + mod) % mod
		}
		sol[i] = ans
	}
	return sol[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
