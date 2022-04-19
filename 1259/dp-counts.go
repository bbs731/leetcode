package leetcode

/*
一道，纯数学的题目， 找到递推的式子， 用 DP 的方法老解决

考虑 0 对  1， 2， 3， ..... n-1  的握手的可能。

(0,1) (2,n-1) +  (0,3) (4, n-1) +  (0,5) (6, n-1)

so let dp[n] 代表 n 个人的握手方案

f[n] = f[0]*f[n-2] + f[2]*f[n-4] + f[4]*f[n-6] +...... + f[n-2]*f[0]
 */
func numberOfWays(numPeople int) int {

	mod := 1000000007
	dp := make([]int, numPeople+1)
	dp[0] = 1
	dp[2] = 1

	for i := 4; i <= numPeople; i = i + 2 {
		for k := 0; k <= i-2; k = k + 2 {
			dp[i] += dp[k] * dp[i-k-2]
			dp[i] %= mod
		}
	}
	return dp[numPeople]

}
