package leetcode

func makeConnected(n int, connections [][]int) int {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		// find with 路径压缩
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	redudent := 0
	for i := 0; i < len(connections); i++ {
		conn := connections[i]
		// do the union
		pa := find(conn[0])
		pb := find(conn[1])
		if pa == pb {
			redudent++
		} else {
			//union
			parent[pa] = pb
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			count++
		}
	}

	if redudent >= count-1 {
		return count - 1
	}
	return -1
}
