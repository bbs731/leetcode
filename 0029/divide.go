package leetcode

import "math"

func divide(dividend int, divisor int) int {
	sign := 0 // positive

	if dividend < 0 {
		dividend = -dividend
		sign = 1
	}

	if divisor < 0 {
		divisor = - divisor
		if sign == 1 {
			sign = 0
		} else {
			sign = 1
		}
	}
	ans := 0

	for dividend >= divisor {
		times := 1
		div := divisor
		for dividend >= div {
			div = div << 1
			times = times << 1
		}

		div >>= 1
		times >>= 1
		ans += times
		dividend -= div
	}

	if sign == 1 {
		if -ans < math.MinInt32 {
			return 1<<31 - 1
		} else {
			return -ans
		}
	}
	if ans > math.MaxInt32 {
		return 1<<31 - 1
	} else {
		return ans
	}
}
