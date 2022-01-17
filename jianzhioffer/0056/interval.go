package offer

import "sort"

type Intervals [][]int

func (t Intervals) Less(i, j int) bool {
	return t[i][0] < t[j][0]
}

func (t Intervals) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Intervals) Len() int {
	return len(t)
}

func merge(intervals [][]int) [][]int {
	var inters Intervals

	inters = intervals
	sort.Sort(inters)

	ans := make([][]int, 0)
	ans = append(ans, inters[0])
	for i := 1; i < len(inters); i++ {
		prev := ans[len(ans)-1]

		if inters[i][0] <= prev[1] {
			// merge
			ans = ans[:len(ans)-1]
			ans = append(ans, []int{prev[0], max(inters[i][1], prev[1])})
		} else {
			ans = append(ans, inters[i])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
