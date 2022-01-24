package leetcode

func lemonadeChange(bills []int) bool {
	charges := [3]int{} // only have 5, 10, 20
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			charges[0]++
		} else if bills[i] == 10 {
			if charges[0] == 0 {
				return false
			}
			charges[0]--
			charges[1]++

		} else {
			//bills[i] = 20
			if charges[1] > 0 {
				charges[1]--
				if charges[0] == 0 {
					return false
				}
				charges[0]--
				charges[2]++
			} else {
				if charges[0] >= 3 {
					charges[0] -= 3
				} else {
					return false
				}
			}
		}
	}
	return true
}
