package leetcode

func removeDuplicates(nums []int) int {
	saved := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[saved] = nums[i]
			saved++
		}
	}
	nums[saved] = nums[len(nums)-1]
	saved++
	return saved
}
