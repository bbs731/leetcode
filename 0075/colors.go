package leetcode

/*
这应该是最好的算法了，
1. 只 scan nums 数组 一次
2. 使用常数内存， O（1） 只有 red, blue, i 这三个计数器

就是太容易写错了
 */
func sortColors(nums []int) {
	red := -1             // 0 stand for red. And red keep current array red end index
	blue := len(nums) - 1 // keep the blue candidateindex

	for i := 0; i <= blue; {
		if nums[i] == 0 {
			// swap  red+1 and i
			nums[red+1], nums[i] = nums[i], nums[red+1]
			red++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			// swap i with blue
			nums[blue], nums[i] = nums[i], nums[blue]
			blue--
		}
	}
}
