package leetcode

func maxScoreSightseeingPair(values []int) int {
	ans, mx := 0, values[0]

	for j := 1; j < len(values); j++ {
		ans = max(ans, mx+values[j]-j)
		mx = max(mx, values[j]+j)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
