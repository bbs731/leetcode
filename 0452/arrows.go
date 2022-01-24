package leetcode

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	ans := 0
	start := points[0]

	isIntersect := func(a, b []int) bool {
		if b[0] <= a[1] {
			return true
		}
		return false
	}
	for i := 1; i < len(points); i++ {
		if isIntersect(start, points[i]) {
			// intersect
			start = []int{points[i][0], min(start[1], points[i][1])} // 取 start, 和 point[i] 最小的 endpoint, 这里容易忽略
		} else {
			ans++
			start = points[i]
		}
	}
	ans++
	return ans
}
