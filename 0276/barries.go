package leetcode

/*
https://leetcode-cn.com/leetbook/read/dynamic-programming-2-plus/5n2tqd/
他是数学的递推关系， 通过 n=1, 2, 3, n-1 总结出来的， 然后把这个递推关系当做 dp 的状态，来求解的，这和传统的 DP 问题不一样

但是官网的数学推倒，还是错的！ 哎！
*/

// f(n) = (k-1) * (f(n-1) + f(n-2))  这种公式 是想不到的！
func numWays(n int, k int) int {

	if n <= 1 {
		return k
	}

	f := make([]int, n)
	f[0] = k
	f[1] = k * k
	for i := 2; i < n; i++ {
		f[i] = k*f[i-1] - f[i-2]
	}
	return f[n-1]
}
