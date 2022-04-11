package leetcode

func countNumbersWithUniqueDigits(n int) int {

	counts := 10

	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}

	for i := 2; i <= n; i++ {
		counts += getNum(i)
	}

	return counts
}

func getNum(n int) int {

	/*
	9 * 9 * 8 * 7....
	 */
	nums := 9
	choice := 9

	for n > 1 {
		nums *= choice
		choice--
		n--
	}
	return nums
}
