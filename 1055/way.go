package leetcode

func shortestWay(source string, target string) int {
	cnt := 0

	for i := 0; i < len(target); {
		match := false
		// try to match greedy
		for j := 0; j < len(source) && i < len(target); {
			if source[j] == target[i] {
				match = true
				j++
				i++
			} else {
				j++
			}
		}

		if !match {
			return -1
		}
		cnt++
	}
	return cnt
}
