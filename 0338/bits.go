package leetcode

func count(n int) int {
	cnt := 0
	for n > 0 {
		if n&1 > 0 {
			cnt++
		}
		n = n >> 1
	}
	return cnt
}
func countBits(n int) []int {

	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = count(i)
	}
	return ans
}
