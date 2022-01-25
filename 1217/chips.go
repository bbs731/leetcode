package leetcode

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for i := 0; i < len(position); i++ {
		if position[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if even < odd {
		return even
	}
	return odd
}
