package leetcode

/*
这题是在受虐吗？ 看官方的解题：真是绝了啊！
 */
func getSum(a int, b int) int {
	carry := 0
	var mask, ans, ua, ub uint64
	mask = 1
	ans = 0

	ua = uint64(a)
	ub = uint64(b)

	for ua != 0 || ub != 0 || carry != 0 {
		ab := ua & mask
		bb := ub & mask

		if carry > 0 {
			if ab > 0 {
				if bb > 0 {
					carry = 1
					ans = ans | mask
				} else {
					// bb == 0
					carry = 1
				}
			} else {
				// ab == 0
				if bb > 0 {
					carry = 1
				} else {
					carry = 0
					ans = ans | mask
				}
			}

		} else { // carry == 0
			if ab > 0 {
				if bb > 0 {
					carry = 1
				} else {
					ans = ans | mask
				}
			} else {
				if bb > 0 {
					ans = ans | mask
				} else {
					// do nothing
				}
			}
		}

		// a b clear the mask
		ua = ua & ^mask
		ub = ub & ^mask
		if mask == 1<<63 {
			break
		} else {
			mask <<= 1
		}
	}

	return int(ans)
}
