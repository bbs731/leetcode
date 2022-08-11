package leetcode

import (
	"sort"
	"strings"
)

/*
欧拉通路，回路问题
https://oi-wiki.org/graph/euler/


时间复杂度  O(n + mlogm)
涉及到的问题：
1. 图的存储，  链式向前星，不适合 edge 的删除操作。 如何存储图， 并可以有效的删除边 edge
2. 如何判断存在欧拉通路或者回路，对于有向图和无向图。
3. 如何选取起始点， 对于本题不存在这个问题。 start from JFK
 */

type edge struct {
	to     string
	exists bool
}

func findItinerary(tickets [][]string) []string {

	beg := make(map[string][]edge)

	for i := 0; i < len(tickets); i++ {
		beg[tickets[i][0]] = append(beg[tickets[i][0]], edge{tickets[i][1], true})
	}

	for _, v := range beg {
		sort.Slice(v, func(i, j int) bool {
			res := strings.Compare(v[i].to, v[j].to)
			if res <= 0 {
				return true
			}
			return false
		})
	}

	ans := make([]string, 0)

	var Hierholzer func(string)
	Hierholzer = func(x string) {

		for i := 0; i < len(beg[x]); i++ {
			if beg[x][i].exists {
				beg[x][i].exists = false // meaning remove the edge
				Hierholzer(beg[x][i].to)
			}
		}
		ans = append(ans, x)
	}

	Hierholzer("JFK")

	// reverse the ans
	i, j := 0, len(ans)-1
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return ans
}
