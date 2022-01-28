package leetcode

func regionsBySlashes(grid []string) int {
	n := len(grid[0]) // we will have n x n grid
	// we will made 1x1 grid two components, so we will have total 2*n*n 个 components, let's component start with index 1
	//这里最主要的问题，就是， parse grid []string 到一个强联通的表达形式

	N := 2 * n * n
	parents := make([]int, N+1)
	for i := 1; i <= N; i++ {
		parents[i] = i
	}

	grids := make([][]int, 0) // we will parse grid []string into grids representation means grids[i][0] is connected with grids[i][1]

	for i := 0; i < n; i++ {
		row := grid[i]
		for j := 0; j < n; j++ {
			itema := 2*n*i + 2*j + 1 // current two item's index
			itemb := 2*n*i + 2*j + 2

			if row[j] == ' ' {
				grids = append(grids, []int{itema, itemb})
			}
			// 处理上面
			if i > 0 {
				upa := 2*n*(i-1) + 2*j + 1
				upb := 2*n*(i-1) + 2*j + 2

				if row[j] == ' ' || row[j] == '/' {
					if grid[i-1][j] == ' ' {
						grids = append(grids, []int{upa, itema})
					} else if grid[i-1][j] == '/' {
						grids = append(grids, []int{upb, itema})
					} else {
						grids = append(grids, []int{upa, itema})
					}
				} else if row[j] == '\\' {
					if grid[i-1][j] == ' ' {
						grids = append(grids, []int{upa, itemb})
					} else if grid[i-1][j] == '/' {
						grids = append(grids, []int{upb, itemb})
					} else {
						grids = append(grids, []int{upa, itemb})
					}
				}
			}

			if j > 0 { // 处理左面 , 左边比较单一， 无论左边是什么， 都可以 join leftb, 自行验证吧！
				//lefta := 2*n*i + 2*(j-1) + 1
				leftb := 2*n*i + 2*(j-1) + 2
				grids = append(grids, []int{leftb, itema})

			}
		}
	}
	// calculate connected components in grids
	var find func(int) int
	find = func(x int) int {
		if parents[x] != x {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}

	for i := 0; i < len(grids); i++ {
		//union
		parents[find(grids[i][0])] = parents[find(grids[i][1])]

	}
	// check how many connected components
	components := 0
	for i := 1; i <= N; i++ {
		if parents[i] == i {
			components++
		}
	}
	return components
}
