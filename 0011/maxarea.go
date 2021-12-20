package leetcode

/*
思路：
指针对撞的问题

性质： height[start]  height[end], 保留两个值当中高的那个, 然后 update start 或者 end。
*/
func maxArea(height []int) int {
	var area, candidate int
	start, end := 0, len(height)-1

	for start < end {
		if height[start] < height[end] {
			candidate = height[start] * (end - start)
			start = start + 1
		} else if height[start] > height[end] {
			candidate = height[end] * (end - start)
			end = end - 1
		} else {
			candidate = height[start] * (end - start)
			start = start + 1
			end = end - 1
		}

		if candidate > area {
			area = candidate
		}
	}
	return area
}
