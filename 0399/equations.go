package leetcode

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	index := 0
	tables := make(map[string]int) // keep the variables index
	neighbours := make(map[int]map[int]struct{})
	results := make(map[[2]int]float64)

	for i := 0; i < len(equations); i++ {
		equation := equations[i]
		if _, ok := tables[equation[0]]; !ok {
			tables[equation[0]] = index
			index++
		}
		if _, ok := tables[equation[1]]; !ok {
			tables[equation[1]] = index
			index++
		}

		if _, ok := neighbours[tables[equation[0]]]; !ok {
			neighbours[tables[equation[0]]] = make(map[int]struct{})
		}
		// add into adj list for each other
		neighbours[tables[equation[0]]][tables[equation[1]]] = struct{}{}
		results[[2]int{tables[equation[0]], tables[equation[1]]}] = values[i]

		if _, ok := neighbours[tables[equation[1]]]; !ok {
			neighbours[tables[equation[1]]] = make(map[int]struct{})
		}
		neighbours[tables[equation[1]]][tables[equation[0]]] = struct{}{}
		results[[2]int{tables[equation[1]], tables[equation[0]]}] = 1 / values[i]
	}

	//stores := make([]float64, n) // keep the final variables float value
	stores := make(map[int]float64)
	stores_components := make(map[int]int)

	var dfs func(int, int)
	dfs = func(start int, components int) {
		// we have invariants that start's value is known
		for k := range neighbours[start] {
			// if k is not visited
			if _, ok := stores[k]; !ok {
				// cal the k's float value
				// at the same time k is mark as visited now
				stores[k] = stores[start] / results[[2]int{start, k}]
				stores_components[k] = components
				dfs(k, components)
			}
		}
	}

	// 我们用 dfs 来找强联通量
	components := 0
	for i := 0; i < index; i++ {
		if _, ok := stores[i]; !ok {
			components++
			stores[i] = 1
			stores_components[i] = components
			dfs(i, components)
		}
	}

	ans := make([]float64, 0)
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		n, ok := tables[query[0]]
		if !ok {
			ans = append(ans, -1)
			continue
		}
		d, ok := tables[query[1]]
		if !ok {
			ans = append(ans, -1)
			continue
		}
		if stores_components[n] != stores_components[d] {
			// 如果来自于不同的 strong connected components 也返回 -1
			ans = append(ans, -1)
			continue
		}
		ans = append(ans, stores[n]/stores[d])
	}
	return ans
}
