package leetcode

/*
[1, 2, 2]
 */

/* stupid answer
自己检讨一下

目前是遇到这种题， 答不上来
*/

func recursive(values []int) int {
	if len(values) < 2 {
		return 0
	}

	N := len(values)
	ans := 0

	start := 0
	end := start + 1
	val := values[end] - end
	// now we have start, lets found end
	for i := start + 1; i < N; i++ {

		if values[i]-i > val {
			val = values[i] - i
			end = i
		}
	}

	// now we need to walk backward
	val = values[end-1] + end - 1
	start = end - 1
	for i := end - 1; i >= 0; i-- {
		if values[i]+i > val {
			val = values[i] + i
			start = i
		}
	}

	ans = values[start] + values[end] + start - end
	return max(ans, recursive(values[end:]))
}

func maxScoreSightseeingPair(values []int) int {
	return recursive(values)
}
