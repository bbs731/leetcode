package leetcode

import (
	"math"
)

/*
 非常简单的逻辑， 为何，最后，写的如此的复杂啊！
 */
func minSumOfLengths(arr []int, target int) int {

	const INFINITY = math.MaxInt64 >> 1
	ans := INFINITY

	n := len(arr)
	prefix := make([]int, n+1)
	presave := make(map[int]int)
	prerecord := make([]int, n+1)
	presave[0] = 0
	prerecord[0] = INFINITY

	reverse := make([]int, n+1)
	reversesave := make(map[int]int)
	reverserecord := make([]int, n+1)
	reverserecord[n] = INFINITY

	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
		presave[prefix[i]] = i
		prerecord[i] = prerecord[i-1]
		if v, ok := presave[prefix[i]-target]; ok {
			prerecord[i] = min(prerecord[i], i-v)
		}
	}

	reversesave[0] = n
	for i := n - 1; i >= 0; i-- {
		reverse[i] = reverse[i+1] + arr[i]
		reversesave[reverse[i]] = i
		reverserecord[i] = reverserecord[i+1]

		if v, ok := reversesave[reverse[i]-target]; ok {
			reverserecord[i] = min(reverserecord[i], v-i)
		}
	}

	for i := 0; i < n; i++ {
		if prerecord[i] != INFINITY && reverserecord[i] != INFINITY {
			ans = min(ans, prerecord[i]+reverserecord[i])
		}
	}

	if ans == INFINITY {
		return -1
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
