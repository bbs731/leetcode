package leetcode

import (
	"strconv"
	"strings"
)

/*
很简单的一道题，
但是，考虑到你现在的岁数， 逻辑上还是很容易出错的, 所以仔细考虑
 */
func dayOfYear(date string) int {
	subs := strings.Split(date, "-")
	year, _ := strconv.Atoi(subs[0])
	month, _ := strconv.Atoi(subs[1])
	day, _ := strconv.Atoi(subs[2])

	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	days := 0

	for i := 1; i < month; i++ {
		days += months[i-1]
	}
	days += day
	if (year%400 == 0 || (year%4 == 0 && year%100 != 0)) && month > 2 {
		days += 1
	}
	return days
}
