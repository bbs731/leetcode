package leetcode

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

/*
binary search template II
 */
func firstBadVersion(n int) int {
	i, j := 1, n

	for i < j {
		h := i + (j-i)>>1
		if isBadVersion(h) {
			j = h
		} else {
			i = h + 1
		}
	}
	return j
}
