package leetcode

import (
	"sort"
)

/*
	最后变成了， 一个预处理 O(mn) 之后， 二分查找的问题了！
	这里我看不出来DP
 */

func minWindow(s1 string, s2 string) string {
	n := len(s1)
	m := len(s2)

	pos := make(map[byte][]int)

	for p := 0; p < len(s2); p++ {
		pos[s2[p]] = make([]int, 0)
	}
	// preprocessing

	for i := 0; i < n; i++ {
		if _, ok := pos[s1[i]]; ok {
			pos[s1[i]] = append(pos[s1[i]], i)
		}
	}

	found := false
	ans := s1

	for _, v := range pos[s2[m-1]] {
		f, largestStart := findLargestBeginIndex(pos, s1, s2, v)
		if f {
			found = true
			if v-largestStart+1 < len(ans) {
				ans = s1[largestStart : v+1]
				//fmt.Println(ans)
			}
		}
	}
	//for k, v := range pos {
	//	fmt.Printf("%c, %v\n", k, v)
	//}
	if !found {
		return ""
	}
	return ans
}

func findLargestBeginIndex(pos map[byte][]int, s1 string, pattern string, endpos int) (bool, int) {
	i := len(pattern) - 2

	for i >= 0 {
		index := sort.SearchInts(pos[pattern[i]], endpos)
		if index == 0 {
			if i == 0 && len(pos[pattern[i]]) != 0 && pos[pattern[i]][0] < endpos {
				return true, pos[pattern[i]][0]
			}
			return false, -1
		}
		endpos = pos[pattern[i]][index-1]
		i--
	}
	return true, endpos
}
