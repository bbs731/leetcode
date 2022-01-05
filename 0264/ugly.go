package leetcode

var uglies map[int]bool

func isUgly(n int) bool {
	if n&0x01 != 1 {
		if _, ok := uglies[n/2]; ok {
			return true
		}
		return false
	}
	if n/5*5 == n {
		if _, ok := uglies[n/5]; ok {
			return true
		}
		return false
	}
	if n/3*3 == n {
		if _, ok := uglies[n/3]; ok {
			return true
		}
		return false
	}

	return false
}

func nthUglyNumber(n int) int {
	found := 1
	number := 1
	ans := 1
	uglies = make(map[int]bool)
	uglies[1] = true

	for found != n {
		number++
		if isUgly(number) {
			found++
			ans = number
			uglies[ans] = true
		}
	}
	return ans
}
