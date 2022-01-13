package leetcode

func largestAltitude(gain []int) int {
	ans := 0

	t := 0
	for i := 0; i < len(gain); i++ {
		t += gain[i]
		ans = max(ans, t)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
