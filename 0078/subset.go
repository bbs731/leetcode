package leetcode

/*
backtracking 的题， 一定会卡死你的，多练吧!

在强调一遍， subset 会搞死你！
*/

func generate(nums []int, k int) [][]int {
	ans := make([][]int, 0)
	if k == 0 {
		ans = append(ans, []int{})
		ans = append(ans, []int{nums[0]})
		return ans
	}

	subs := generate(nums, k-1)
	ans = append(ans, subs...)
	for i := 0; i < len(subs); i++ {
		s := []int{}
		s = append(s, subs[i]...)
		s = append(s, nums[k])
		ans = append(ans, s)
	}
	return ans
}
func subsets(nums []int) [][]int {
	return generate(nums, len(nums)-1)
}
