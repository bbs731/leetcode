package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
边界条件
[0]
[0,1]

 */
func canJump(nums []int) bool {

	far := 0
	for i := 0; i < len(nums); i++ {
		far = max(far, i+nums[i])
		if i == far {
			break
		}
	}
	return far >= len(nums)-1 /* 这里最重要 len(nums) 还是 len(nums) -1 ?  */
}
