package leetcode


/*
 wow! 这道题，很有欺骗性， 顺序不一样，算不同的序列。

这道题，给我们的经验，就是， 首先看出来是 一定是个DP 问题。

最重要的一点， DP 的计算顺序，不同，计算结果不同。  先按照 nums[i] 一个一个number 的作为第一次 loop 算出来的结果，是不考虑顺序的。
所以，第一次loop 应该按照 1,2,3, target 来 loop

 */
func combinationSum4(nums []int, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1

	getter := func(i int) int {
		if i < 0 {
			return 0
		}
		return dp[i]
	}

	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			dp[j] += getter(j - nums[i])
		}
	}
	return dp[target]
}
