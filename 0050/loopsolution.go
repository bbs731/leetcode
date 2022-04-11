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

/*

using ll = long long;
const int M = 2;

struct Ma
{
    int a[M][M];
    Ma()
    {
        memset(a, 0, sizeof(a));
    }

    void init() // 复位为单位阵
    {
        a[0][0] = a[1][1] = 1;
        a[0][1] = a[1][1] = 0;
    }

    Ma operator*(const Ma& B) const
    {
        Ma ans;
        for(int i = 0; i < M; ++i)
            for(int j = 0; j < M; ++j)
                for(int k = 0; k < M; ++k)
                    ans.a[i][j] += a[i][k] * B.a[k][j];
        return ans;
    }

    Ma operator^(int n) const
    {
        Ma ans;
        ans.init();
        Ma A = *this; // 拷贝一个出来用于自乘
        while(n)
        {
            if(n & 1)
                ans = ans * A;
            A = A * A;
            n >>= 1;
        }
        return ans;
    }
};


 */
