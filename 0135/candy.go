package leetcode

import "math"


/*
思路， 想的，还是有点复杂了！

找时间在做一遍吧！换个思路
 */
func candy(ratings []int) int {

	candies := make([]int, len(ratings))
	// 找到最小的那个点， 然后总那个点，想两边发散

	smallest := math.MaxInt64
	si := 0
	for i := 0; i < len(ratings); i++ {
		if ratings[i] < smallest {
			smallest = ratings[i]
			si = i
		}
	}

	candies[si] = 1
	stack := make([]int, 0)
	stack = append(stack, si)

	for len(stack) > 0 {
		child := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// walk from child to left
		for i := child - 1; i >= 0; i-- {
			if candies[i] != 0 { //visited i before
				if ratings[i] > ratings[i+1] {
					candies[i] = max(candies[i], candies[i+1]+1)
				} else {
					break
				}

			} else {
				if ratings[i] > ratings[i+1] {
					candies[i] = candies[i+1] + 1
				} else if ratings[i] == ratings[i+1] {
					// special treatment for equal case
					// and maybe update later by visited case
					candies[i] = 1
				} else {
					//ratings[i] < ratings[i+1]
					j := i
					for ; j >= 0 && ratings[j] <= ratings[j+1]; j-- {

					}
					j = j + 1

					if candies[j] == 0 {
						// not visited j before
						candies[j] = 1
						stack = append(stack, j)
					}
					// end of left walk
					break
				}
			}
		}

		// walk from child to right
		for i := child + 1; i < len(ratings); i++ {
			if candies[i] != 0 { //visited before
				if ratings[i] > ratings[i-1] {
					candies[i] = max(candies[i], candies[i-1]+1)
				} else {
					break
				}

			} else {
				if ratings[i] > ratings[i-1] {
					candies[i] = candies[i-1] + 1
				} else if ratings[i] == ratings[i-1] {
					//special equal case
					candies[i] = 1
				} else {
					//ratings[i] < ratings[i-1]
					j := i
					for ; j < len(ratings) && ratings[j] <= ratings[j-1]; j++ {

					}
					j = j - 1
					if candies[j] == 0 {
						candies[j] = 1
						stack = append(stack, j)
					}
					// end of right walk
					break
				}
			}
		}
	}

	ans := 0

	for i := 0; i < len(candies); i++ {
		ans += candies[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
