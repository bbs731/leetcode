package leetcode

func monotoneIncreasingDigits(n int) int {
	digits := make([]int, 0)

	num := n
	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	keep := -1
	// digits is in reverse order, e.g  12319, digits= [9,1,3,2,1]
	for i := 1; i < len(digits); {
		if digits[i] <= digits[i-1] {
			i++
			continue
		}
		// otherwise we have a increase here
		//digits[i] > digits[i-1]
		digits[i] -= 1
		digits[i-1] = 9
		keep = i - 1
	}

	ans := 0
	for i := len(digits) - 1; i >= 0; i-- {
		ans *= 10
		if i <= keep {
			ans += 9
		} else {
			ans += digits[i]
		}
	}
	return ans
}
