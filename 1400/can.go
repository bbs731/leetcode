package leetcode

func canConstruct(s string, k int) bool {
	table := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		table[s[i]]++
	}

	even, odd := 0, 0

	for _, v := range table {
		if v%2 == 0 {
			even += v
		} else {
			odd++
			even += v - 1
		}
	}

	if k >= odd && k <= odd+even {
		return true
	}
	return false
}
