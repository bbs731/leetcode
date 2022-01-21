package leetcode

var ans []string

func generate(table map[string][]string, digits string, k int, content string) {
	if k == len(digits) {
		ans = append(ans, content)
		return
	}
	// fill up the kth element

	candidates := table[string(digits[k])]
	for i := 0; i < len(candidates); i++ {
		generate(table, digits, k+1, content+candidates[i])
	}
}

func letterCombinations(digits string) []string {

	table := make(map[string][]string)
	table["2"] = []string{"a", "b", "c"}
	table["3"] = []string{"d", "e", "f"}
	table["4"] = []string{"g", "h", "i"}
	table["5"] = []string{"j", "k", "l"}
	table["6"] = []string{"m", "n", "o"}
	table["7"] = []string{"p", "q", "r", "s"}
	table["8"] = []string{"t", "u", "v"}
	table["9"] = []string{"w", "x", "y", "z"}

	ans = make([]string, 0)
	if len(digits) == 0 {
		return ans
	}
	generate(table, digits, 0, "")
	return ans
}
