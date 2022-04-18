package leetcode

func minEatingSpeed(piles []int, h int) int {
	sum := 0
	largest := piles[0]

	for i := 0; i < len(piles); i++ {
		sum += piles[i]
		largest = max(largest, piles[i])
	}

	l, r := max(sum/h, 1), largest
	for l < r {
		total := 0
		mid := (r + l) >> 1

		for i := 0; i < len(piles); i++ {
			if piles[i]%mid == 0 {
				total += piles[i] / mid
			} else {
				total += piles[i]/mid + 1
			}
		}
		/*
		 	这里，就是你的梦魇！
		 */

		if total <= h {
			r = mid
		} else if total > h {
			l = mid + 1
		}
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
