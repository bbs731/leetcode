package leetcode

import (
	"sort"
)

/*
好题啊！ 你浪费了一个练习  DP  的 好机会！
 */
type record struct {
	endtime       int
	originalindex int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	records := []*record{}

	dp := make([]int, n+1)

	for i := 0; i < n; i++ {
		records = append(records, &record{endTime[i], i})
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].endtime < records[j].endtime
	})

	ans := 0

	//dp[0] = profit[records[0].originalindex]  // 直接初始化，结果是错的。
	for i := 1; i <= n; i++ {
		//pos := upperbound(records, startTime[records[i-1].originalindex])
		pos := sort.Search(n, func(mid int) bool {
			return records[mid].endtime > startTime[records[i-1].originalindex]
		})
		// 注意， dp[] 的 starting index 是 1 不是 0 
		dp[i] = max(dp[i-1], dp[pos]+profit[records[i-1].originalindex])

		ans = max(ans, dp[i])
	}
	return ans
}

// 这是 二分查找中 upperbound 的写法！
func upperbound(records []*record, starttime int) int {
	n := len(records)

	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if records[h].endtime <= starttime {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
