package leetcode

/*

真是走狗屎运了， 一遍就过了

有 O(n) 的解法， 看官方解题的答案

with current index as left, find the right pillar
// 1. 找到 current index 右边第一个比 current index pillar 相等或者高的
// 2. 如果 step 1找不到, 说明右边的所有 pillar 都不 current index pillar 矮， 这时候需要返回， 所有右边 pillar 中，最高的，
// 如果有多个相等的， 取 index 最靠右边的
*/
func findNexPillar(left int, height []int) int {

	N := len(height)
	index := left
	max := 0
	for i := left + 1; i < N; i++ {
		if height[i] >= height[left] {
			index = i
			break
		}

		if height[i] >= max {
			max = height[i]
			index = i
		}
	}
	return index
}

func trap(height []int) int {

	N := len(height)
	water := 0
	left := 0

	for left < N-1 {
		right := findNexPillar(left, height)
		h := min(height[left], height[right])

		for i := left + 1; i < right; i++ {
			water += h - height[i]
		}
		left = right
	}
	return water
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
