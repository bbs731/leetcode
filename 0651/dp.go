package leetcode

func maxA(n int) int {

	f := make([]int, max(n+1, 4))
	f[0] = 0
	f[1] = 1
	f[2] = 2
	f[3] = 3

	for i := 4; i <= n; i++ {
		f[i] = f[i-1] + 1
		for k := 3; i-k > 0; k++ {
			f[i] = max(f[i], f[i-k]*(k-1))
		}
	}
	return f[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
