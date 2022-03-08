package leetcode

import "sort"

/*
Wow !  用两天的时间， 创造了一堆屎一样的代码！

需要总结的地方，  sort.SliceStable 和 sort.Slice 是有区别的。 看 test.main.go

比所找的数大的最小值， 和比这个数小的里面的最大值。 这个如果有一个 balanced binary search tree, 是不是就很好解决了？
看看 java 里面的 TreeMap 是如何实现的？
 */
type record struct {
	val   int
	index int
}

func oddEvenJumps(arr []int) int {

	n := len(arr)
	records := make([]record, n)
	records2 := make([]record, n)
	for i := 0; i < n; i++ {
		records[i] = record{arr[i], i}
		records2[i] = record{arr[i], i}
	}
	sort.SliceStable(records, func(i, j int) bool {
		if records[i].val < records[j].val {
			return true
		} else if records[i].val > records[j].val {
			return false
		}
		return i < j
	})
	sort.SliceStable(records2, func(i, j int) bool {
		if records2[i].val < records2[j].val {
			return true
		} else if records2[i].val > records2[j].val {
			return false
		}
		return i > j
	})

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, 2) // 0 stands for event 1 stands for odd
	}

	dp[n-1][0] = true
	dp[n-1][1] = true

	for i := n - 2; i >= 0; i-- {
		// odd jumps
		// search for record{arr[i], i} pos
		pos := sort.Search(n, func(p int) bool {
			return records[p].val >= arr[i]
		})

		for j := pos; j < n; j++ {
			if records[j].index > i {
				dp[i][1] = dp[records[j].index][0]
				break
			}
		}
		pos2 := sort.Search(n, func(p int) bool {
			return records2[p].val >= arr[i]
		})

		val := records2[pos2].val
		// walk to the right if encounter equal values
		for pos2+1 < n && records2[pos2+1].val == val {
			pos2++
		}

		for j := pos2 - 1; j >= 0; j-- {
			if records2[j].index > i {
				dp[i][0] = dp[records2[j].index][1]
				break
			}
		}
	}
	cnts := 0
	for i := 0; i < n; i++ {
		if dp[i][1] == true {
			cnts++
		}
	}
	return cnts
}
