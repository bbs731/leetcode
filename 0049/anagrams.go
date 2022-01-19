package leetcode

import "sort"

/*
golang 里面， string to []byte, []byte to string 的转换， 怎么就是记不住呢？
https://segmentfault.com/a/1190000037679588
 */
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
