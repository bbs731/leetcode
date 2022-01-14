package leetcode

/*
最标准的解法， 利用 bit 运算
 */

func singleNumber(nums []int) int {

	simple := 0

	for i := 0; i < len(nums); i++ {
		simple ^= nums[i]
	}

	return simple
}
