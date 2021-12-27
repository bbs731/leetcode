package leetcode

func findMiddleIndex(nums []int) int {

	sum := make([]int, len(nums))

	sum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}

	if sum[len(nums)-1]-sum[0] == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if sum[i-1] == sum[len(nums)-1]-sum[i] {
			return i
		}
	}
	return -1
}
