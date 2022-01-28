package leetcode

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	parent := make([]int, n)
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

	// treat pairs 找到连通量， 连通量都是string index position
	for i := 0; i < len(pairs); i++ {
		pair := pairs[i]
		// union and find
		parent[find(pair[0])] = find(pair[1])
	}

	/*
		这里的技巧: 统计，每一个连通量 都是由哪些index组成的
	 */
	sets := make(map[int][]int)
	for i := 0; i < n; i++ {
		p := find(i)
		if _, ok := sets[p]; !ok {
			sets[p] = []int{i}
		} else {
			sets[p] = append(sets[p], i)
		}
	}

	ans := make([]byte, n)
	for _, v := range sets {
		index := make([]int, len(v))
		copy(index, v)
		sort.Slice(index, func(i, j int) bool { return index[i] < index[j] })
		sort.Slice(v, func(i, j int) bool { return s[v[i]] < s[v[j]] })
		for j := 0; j < len(v); j++ {
			ans[index[j]] = s[v[j]]
		}
	}

	return string(ans)
}
