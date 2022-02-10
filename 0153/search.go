package leetcode
/*
binary search Template II
 */
func findMin(nums []int) int {
	i, j := 0, len(nums)
	last := nums[len(nums)-1]

	for i < j {
		h := i + (j-i)>>1
		if nums[h] > last {
			i = h + 1
		} else {
			j = h
		}
	}
	return nums[j]
}
