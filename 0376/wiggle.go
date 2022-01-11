package leetcode

/* wow !
这个答案是错了， O(n^2)

有 O（n)的时间复杂度的解!
重新找状态方程吧，大哥！
*/


func wiggleMaxLength(nums []int) int {
	N := len(nums)
	prev := make([]int, N)
	direction := make([]int, N) // 0 stands for unknown, 1 stands for increase from prev element, -1 stands for decrease from prev statement
	length := make([]int, N)

	prev[0] = -1
	direction[0] = 0
	length[0] = 1

	ans := 1
	for i := 1; i < N; i++ {
		prev[i] = -1
		direction[i] = 0
		length[i] = 1
		maxlength := 0
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && direction[j] != 1 {
				if length[j]+1 > maxlength {
					maxlength = length[j] + 1
					prev[i] = j
					direction[i] = 1
					length[i] = maxlength
				}
			} else if nums[i] < nums[j] && direction[j] != -1 {
				if length[j]+1 > maxlength {
					maxlength = length[j] + 1
					prev[i] = j
					direction[i] = -1
					length[i] = maxlength
				}
			} // nums[i] == nums[j] 这种情况， 在初始化 prev[i], direction[i], length[i] 的时候，已经 cover 住了
		}
		ans = max(ans, maxlength)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
