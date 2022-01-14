package leetcode

/*
Boyer-Moore 算法 可以， 把 stack 的空间， 变成  O(1)， 用一个 counter int, 和 candidate int 两个数来代替
 */
func majorityElement(nums []int) int {

	stack := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if len(stack) > 0 && stack[len(stack)-1] != nums[i] {
			stack = stack[:len(stack)-1] // pop
			continue
		}
		stack = append(stack, nums[i])
	}

	return stack[len(stack)-1]
}
