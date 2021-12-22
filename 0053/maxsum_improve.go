package leetcode

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	sum := 0

	for _, n := range nums {
		sum += n
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
