package leetcode

// n 个城市

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parent := make(map[int]int)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	components := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if isConnected[i][j] == 1 {
				pi := find(i)
				pj := find(j)
				parent[pi] = pj
			}
		}
	}

	//keep := make(map[int]struct{})
	//for i := 0; i < n; i++ {
	//	pi := find(i)
	//	if _, ok := keep[pi]; !ok {
	//		components++
	//		keep[pi] = struct{}{}
	//	}
	//}

	// 还有个技巧在这里, count, how many sets in total
	for i := 0; i < n; i++ {
		if find(i) == i {
			components++
		}
	}

	return components
}
