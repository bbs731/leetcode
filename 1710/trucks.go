package leetcode

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]  // 这里需要倒序
	})

	count := 0
	remain := truckSize

	for i := 0; i < len(boxTypes); i++ {
		if boxTypes[i][0] <= remain {
			count += boxTypes[i][0] * boxTypes[i][1]
			remain -= boxTypes[i][0]
		} else {
			count += remain * boxTypes[i][1]
			break
		}
	}
	return count
}
