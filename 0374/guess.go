package leetcode

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {

	i, j := 1, n

	for i <= j {
		h := i + (j-i)>>1
		g := guess(h)
		if g == 0 {
			return h
		} else if g == 1 {
			i = h + 1
		} else {
			j = h - 1
		}
	}
	return -1
}
