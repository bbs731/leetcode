package main

import (
	"fmt"
	"sort"
)

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

	i := 3
	pos := sort.Search(n, func(p int) bool {
		return records[p].val >= arr[i]
	})
	fmt.Println(pos)
	fmt.Println(records)
	fmt.Println(records2)

	pos2 := sort.Search(n, func(p int) bool {
		return records2[p].val >= arr[i]
	})
	fmt.Println(pos2)
	for i := n - 2; i >= 0; i-- {
		// odd jumps
		// search for record{arr[i], i} pos
		//pos := sort.Search(n, func(p int) bool {
		//	return records[p].val >= arr[i] && records[p].index >= i
		//})
		//
		//
		//pos2 := sort.Search(n, func(p int) bool {
		//	return records2[p].val >= arr[i] && records2[p].index >= i
		//})
	}

	return 0
}

func main() {
	arrs := []int{2, 3, 1, 1, 4, 2, 3, 1, 1, 4}
	oddEvenJumps(arrs)
}
