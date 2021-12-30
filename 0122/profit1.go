package leetcode

func maxProfit(prices []int) int {
	N := len(prices)

	acc := 0
	min, max := prices[0], prices[0]
	climbing := true
	for i := 1; i < N; i++ {
		if prices[i] >= prices[i-1] {
			//climbing
			max = prices[i]
			climbing = true
		} else {
			if climbing {
				acc += max - min
				climbing = false
			}
			min = prices[i]
		}
	}

	// we do the last for end climbing
	if climbing {
		acc += max - min
	}

	return acc
}
