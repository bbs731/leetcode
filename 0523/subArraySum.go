package leetcode

/*
 又是一道，前缀和 +  Hasatable 的例子。

前缀和，有多种变形， 需要根据题目来进行设计，
譬如， 加 1 减 1, 想加取余数等等
 */
func checkSubarraySum(nums []int, k int) bool {
	pre := 0
	cache := map[int]int{} // cache the index
	cache[0] = -1

	found := false
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		mod := pre % k

		if _, ok := cache[mod]; ok {
			if i-cache[mod] > 1 {
				found = true
			}
		} else {
			cache[mod] = i
		}
	}
	return found
}
