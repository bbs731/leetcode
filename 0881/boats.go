package leetcode

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	left, right := 0, len(people)-1
	counts := 0

	for left <= right {
		if people[left]+people[right] > limit {
			counts++
			right--
		} else {
			counts++
			left++
			right--
		}
	}
	return counts
}
