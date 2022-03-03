package leetcode

import "sort"

/*
[[2,11],[6,16],[11,16]]
[[9,16],[6,16],[1,9],[3,15]]
 */

func minMeetingRooms(intervals [][]int) int {
	// 安装会议结束时间排序 (我被打败了， 按照 开始时间排序）
	sort.Slice(intervals, func(i, j int) bool {
		//if intervals[i][1] < intervals[j][1] {
		//	return true
		//} else if intervals[i][1] == intervals[j][1] {
		//	return intervals[i][0] > intervals[j][0]
		//}
		//return false
		return intervals[i][0] < intervals[j][0]
	})

	ans := 1
	// save current used meet room's endtime, index is the room's number
	endtimes := []int{intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < endtimes[0] {
			ans++
			endtimes = append(endtimes, intervals[i][1])
		} else {
			//pos := sort.SearchInts(endtimes, intervals[i][0])
			//if pos == len(endtimes) {
			//	pos = pos - 1
			//}
			//if endtimes[pos] != intervals[i][0] {
			//	pos = pos - 1
			//}
			//// update meeting room pos's endtime
			//endtimes[pos] = intervals[i][1]
			endtimes[0] = intervals[i][1]
		}
		sort.Ints(endtimes)
	}
	return ans
}
