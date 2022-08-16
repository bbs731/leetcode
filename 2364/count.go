package leetcode

func countBadPairs(nums []int) int64 {
	n := len(nums)
	counts := make(map[int]int)

	for i := 0; i < n; i++ {
		counts[nums[i]-i] += 1
	}

	ans := 0
	for i := 0; i < n; i++ {
		//ans += n - 1 - (counts[nums[i]-i] - 1)
		ans += n - counts[nums[i]-i]
	}
	return int64(ans / 2)
}
