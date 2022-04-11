package leetcode

/*
矩阵的快速幂模板， 时间复杂度可以从 O(n) 下降到  O(lgn)
https://leetcode-cn.com/leetbook/read/dynamic-programming-2-plus/xboi93/

while(n)
{
    if(n & 1)
        ans = ans * A; // 这一步是矩阵的乘法
    A = A * A;
    n >>= 1;
}

 */
func pow(x float64, n int) float64 {
	var ans float64
	ans = 1

	for n > 0 {
		if n&1 > 0 {
			ans *= x
		}
		n >>= 1
		x *= x
	}

	return ans
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
