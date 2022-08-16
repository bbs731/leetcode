package leetcode

func taskSchedulerII(tasks []int, space int) int64 {
	schedule := make(map[int]int)
	currday := 0

	for i := 0; i < len(tasks); i++ {
		currday++

		//if currday >= schedule[tasks[i]] {
		//	// currday  can fish the ith task
		//	// then update the schedule for ith task
		//	schedule[tasks[i]] = currday + space + 1
		//} else {
		//	// we need rest on currday
		//	// and we can fastforward the currday and fish the work on that day
		//	currday = schedule[tasks[i]]
		//	schedule[tasks[i]] = currday + space + 1
		//}

		if currday < schedule[tasks[i]] {
			currday = schedule[tasks[i]]
		}
		schedule[tasks[i]] = currday + space + 1
	}

	return int64(currday)
}
