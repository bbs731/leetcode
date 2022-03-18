package leetcode

func maxSubArrayLen(nums []int, k int) int {

	n := len(nums)
	sums := make([]int, n+1)
	occurs := make(map[int]int)
	occurs[0] = 0

	ans := 0

	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]

		if _, ok := occurs[sums[i]]; !ok {
			occurs[sums[i]] = i
		} else {
			// do nothing, since now index i is increaseing
		}

		if p, ok := occurs[sums[i]-k]; ok {
			ans = max(ans, i-p)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
