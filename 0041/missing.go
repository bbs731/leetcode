package leetcode

/*
「真正」满足时间复杂度为 O(N)O(N) 且空间复杂度为 O(1)O(1) 的算法是不存在的，
但是我们可以退而求其次：利用给定数组中的空间来存储一些状态。也就是说，如果题目给定的数组是不可修改的，那么就不存在满足时空复杂度要求的算法；
但如果我们可以修改给定的数组，那么是存在满足要求的算法的。
*/

// 如果不允许修改数组的话， 这道题，是没有空间 O(1) 时间 O(n) 的方案的。
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		for nums[i] >= 1 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			//swap, 让 nums[i] 放到正确的位置 index 上， 就是  nums[i]-1 的 index 上
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i]-1 != i {
			return i + 1
		}
	}
	return n + 1
}
