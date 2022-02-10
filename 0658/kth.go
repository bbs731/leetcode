package leetcode

import "sort"

func findClosestElements(arr []int, k int, x int) []int {

	left, right := 0, len(arr)-1

	for left+1 < right {
		mid := left + (right-left)>>1

		if arr[mid] == x {
			left = mid
			right = mid + 1
			break
		} else if arr[mid] > x {
			right = mid
		} else {
			left = mid
		}
	}
	ans := make([]int, 0)
	for k > 0 {
		k--
		if left < 0 {
			ans = append(ans, arr[right])
			right++
			continue
		}
		if right >= len(arr) {
			ans = append(ans, arr[left])
			left--
			continue
		}
		if x-arr[left] <= arr[right]-x {
			ans = append(ans, arr[left])
			left--
		} else {
			ans = append(ans, arr[right])
			right++
		}
	}
	sort.Ints(ans)
	return ans
}
