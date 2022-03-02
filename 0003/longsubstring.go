package leetcode

/* 思路：
    i,j 表示目前正在侦测的 substring 的起始和终点的index
	m[byte]int 保存从i到j的字符对应的index

	如果发现 s[j] 和 i到j-1之间的一个字符重复，如 s[pos]
	那么清除 从 i 到 pos 的 m[] 更新 i = pos +1


sliding window 

*/
func lengthOfLongestSubstring(s string) int {

	var i, j, length int

	m := make(map[byte]int)
	for ; j < len(s); j++ {
		//found duplicate byte
		if pos, ok := m[s[j]]; ok {
			//update map, clear entry start from i to pos
			for ; i < pos; i++ {
				delete(m, s[i])
			}
			// update i to pos + 1
			i = pos + 1
		} else {
			// update the longest substring's len if found
			if length < j-i+1 {
				length = j - i + 1
			}
		}

		// update m[s[j]] from pos to j or insert new record with key s[j] and value j
		m[s[j]] = j
	}

	return length
}
