package leetcode

func findMaxLength(nums []int) int {

	pre := 0
	index := make(map[int]int)
	index[0] = -1
	ans := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			pre--
		} else {
			pre++
		}

		if _, ok := index[pre]; ok {
			// we have encounter this pre value before
			//我们可以计算长度

			len := i - (index[pre] + 1) + 1
			ans = max(ans, len)
		} else {
			// 第一次出现， 记录一下
			index[pre] = i
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
