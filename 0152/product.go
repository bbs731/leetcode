package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
[2,3,-2,4] = 6
[2,-5,-2,-4,3] = 24
[-2,0,-1] = 0


这种题，狠狠的打脸了，第一遍逻辑有漏铜。
问题: 如何验证自己的解题逻辑是正确的呢？ 用测试用例反证， 但是这是后验！ 如果保证自己的算法逻辑是正确的？
 */

func maxProduct(nums []int) int {
	ans := nums[0]

	positive := make([]int, len(nums))
	negative := make([]int, len(nums))
	if nums[0] > 0 {
		positive[0] = nums[0]
		negative[0] = 0
	} else {
		positive[0] = 0
		negative[0] = nums[0]
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			positive[i], negative[i] = 0, 0
			ans = max(0, ans)
		} else if nums[i] > 0 {
			positive[i] = max(nums[i], positive[i-1]*nums[i])
			negative[i] = negative[i-1] * nums[i]
		} else {
			positive[i] = negative[i-1] * nums[i]
			negative[i] = positive[i-1] * nums[i]
			if negative[i] == 0 {
				negative[i] = nums[i]
			}
		}
		ans = max(ans, positive[i])
	}

	return ans
}
