package leetcode

import "sort"

type record struct {
	starttime     int
	originalindex int
}


/*
这个方法的复杂度是多少啊？  O(2^n) 吗？
 */
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	records := []*record{}

	accum := make([]int, n)
	for i := 0; i < n; i++ {
		accum[i] = profit[i]
	}

	for i := 0; i < n; i++ {
		records = append(records, &record{startTime[i], i})
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].starttime < records[j].starttime
	})

	var dfs func(et int, total int)

	dfs = func(et int, total int) {
		s := sort.Search(n, func(pos int) bool {
			return records[pos].starttime >= et
		})

		for i := s; i < n; i++ {
			origin := records[i].originalindex
			if accum[origin] < total+profit[origin] {
				accum[origin] = total + profit[origin]
				dfs(endTime[origin], accum[origin])
			}
		}
	}

	for i := 0; i < n; i++ {
		origin := records[i].originalindex
		dfs(endTime[origin], profit[origin])
	}

	largest := 0
	for i := 0; i < n; i++ {
		if accum[i] > largest {
			largest = accum[i]
		}
	}

	return largest
}
