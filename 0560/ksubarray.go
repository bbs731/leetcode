package leetcode
/*

我噻， 这道题，可是经典了！

前缀和  （这是一个新的知识点，之前不了解）

前缀和 +  hashtable 解这道题，  record 不能只记录true or false, 需要记录出现的次数。 要不然这个 test case 过不去
[1, -1, 0]
0

ans: 3

 */
func subarraySum(nums []int, k int) int {
	N := len(nums)

	record := make(map[int]int)
	// 前缀和数组
	// sum[i ... j] = preSum[j+1] - preSum[i]
	preSum := make([]int, N+1)
	preSum[0] = 0
	record[0] = 1
	ans := 0
	for i := 0; i < N; i++ {
		preSum[i+1] = preSum[i] + nums[i]
		sumj := preSum[i+1] - k
		if _, ok := record[sumj]; ok {
			ans += record[sumj]
		}
		record[preSum[i+1]]++
	}
	return ans
}
