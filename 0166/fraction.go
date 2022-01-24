package leetcode

import "strconv"

/*
这道题， corner condition 太多需要考虑的地方了。

多锻炼， 争取一次写对！
 */

func fractionToDecimal(numerator int, denominator int) string {
	ans := ""

	if numerator < 0 {
		if denominator > 0 {
			ans += "-"
		} else {
			denominator = - denominator
		}
		numerator = -numerator
	} else if numerator > 0 {
		if denominator < 0 {
			ans += "-"
			denominator = - denominator
		}
	}

	// now we have positive numerator and denominator
	table := make(map[int]int)
	//table[0] = struct{}{}
	div := numerator / denominator
	r := numerator - div*denominator
	ans += strconv.Itoa(div)
	if r != 0 {
		ans += "."
	}

	for r != 0 {
		if pos, ok := table[r]; ok {
			// do special treatment for ans
			ans += ")"
			ans = ans[:pos] + "(" + ans[pos:]
			break
		}
		if len(ans) == 10000 {
			break
		}
		// record the remainder occur position, later use to add ( ) bracket
		table[r] = len(ans)
		numerator = r * 10
		div = numerator / denominator
		ans += strconv.Itoa(div)
		r = numerator - div*denominator
	}
	return ans
}
