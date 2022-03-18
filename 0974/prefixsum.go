package leetcode

func subarraysDivByK(nums []int, k int) int {
	record := make(map[int]int)
	n := len(nums)
	prefixsum := make([]int, n+1)
	ans := 0
	// 这里的初始化，要特别注意
	record[0] = 1

	for i := 1; i <= n; i++ { // prefixsum 这里 i 最大可以取到 n 太容易忘记了， 这里出错了好多次
		prefixsum[i] = ((prefixsum[i-1]+nums[i-1])%k + k) % k
		if v, ok := record[prefixsum[i]]; ok {
			ans += v
		}
		record[prefixsum[i]]++
	}
	return ans
}
