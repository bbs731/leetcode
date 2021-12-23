package leetcode

func permute(nums []int) [][]int {
	if len(nums) == 0 || len(nums) == 1 {
		return [][]int{nums}
	}

	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		// swap first element with i
		nums[0], nums[i] = nums[i], nums[0]
		for _, s := range permute(nums[1:]) {
			candidate := make([]int, 0, len(nums))
			candidate = append(candidate, nums[0])
			candidate = append(candidate, s...)

			res = append(res, candidate)
		}
		//restore
		nums[0], nums[i] = nums[i], nums[0]
	}
	return res
}

/*
经验教训：

	candidate := make([]int, 0, len(nums)
	candidate[0] = nums[0] will cause panic,  知道为什么吗？


	candidate := make([]int, len(nums))
	candidate[0] = nums[0]
	candidate = append(candidate, s...) 结果是错的， 能看出来是为什么吗？


	哎， 对slice 的操作太不熟悉了！
 */
