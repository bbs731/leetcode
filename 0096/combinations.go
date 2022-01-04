package leetcode

/*
f(n) = f(0)*f(n-1) + f(1)*f(n-2) + f(2)*f(n-3) + ..... + f(n-2)*f(1) + f(n-1)*f(0)
 */
func numTrees(n int) int {

	nums := make([]int, n+1)
	nums[0], nums[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			nums[i] += nums[j] * nums[i-1-j]
		}
	}
	return nums[n]
}
