package leetcode

import "strings"

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	parent := make([]int, 26)

	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(p int) int {
		if parent[p] != p {
			parent[p] = find(parent[p])
		}
		return parent[p]
	}
	for i := 0; i < len(s1); i++ {
		a := s1[i]
		b := s2[i]

		pa := find(int(a - 'a'))
		pb := find(int(b - 'a'))
		if pa != pb {
			// do the union
			if pa < pb {
				parent[pb] = pa
			} else {
				parent[pa] = pb
			}
		}
	}

	convert := &strings.Builder{}

	for i := 0; i < len(baseStr); i++ {
		pa := find(int(baseStr[i] - 'a'))
		convert.WriteByte(byte(pa + 'a'))
	}
	return convert.String()
}
