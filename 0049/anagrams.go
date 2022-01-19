package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {

	tables := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		s := strs[i]
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })

		tables[string(bs)] = append(tables[string(bs)], s)
	}

	var ans [][]string

	for _, v := range tables {
		ans = append(ans, v)
	}
	return ans
}
