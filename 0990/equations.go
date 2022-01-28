package leetcode

func equationsPossible(equations []string) bool {
	parent := make(map[byte]byte)

	var find func(byte) byte

	find = func(x byte) byte {
		if _, ok := parent[x]; !ok {
			parent[x] = x
		}
		if parent[x] != x { // 注意了，这里是用的 if, 不是 for
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	//find = func(x byte) byte {
	//	if _, ok := parent[x]; !ok {
	//		parent[x] = x
	//	}
	//	for parent[x] != x {
	//		parent[x] = parent[parent[x]]
	//		x = parent[x]
	//	}
	//	return parent[x]
	//}
	//
	for i := 0; i < len(equations); i++ {
		// parse  the four byte
		item := equations[i]
		x, y := item[0], item[3]

		px := find(x)
		py := find(y)

		if item[1] == '=' { // == case
			if px != py {
				parent[px] = py
			}
		}
	}
	for i := 0; i < len(equations); i++ {
		item := equations[i]
		x, y := item[0], item[3]
		px, py := find(x), find(y)
		if item[1] == '!' {
			if px == py {
				return false
			}
		}
	}
	return true
}
