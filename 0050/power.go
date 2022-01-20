package leetcode

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}

	res := myPow(x, n/2)
	if n%2 == 0 {
		return res * res
	} else {
		return res * res * x
	}
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / pow(x, -n)
	} else {
		return pow(x, n)
	}
}
