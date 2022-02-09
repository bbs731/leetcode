package leetcode

/* 这道题， 是 线性DP 最简单，最基础的问题
	最大连续子数组之和
*/

func maxSubArray(nums []int) int {
	max := nums[0]
	// 似乎声明一个 sum 数组是没必要的， 应为只用到了 sum[i-1]
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if sum[i-1] <= 0 {
			sum[i] = nums[i]
		} else {
			sum[i] = sum[i-1] + nums[i]
		}
		if sum[i] > max {
			max = sum[i]
		}
	}
	return max
}
