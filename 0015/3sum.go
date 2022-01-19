package leetcode

/*
这不是最好的解法， 可以不使用 hashtable
看官方的解法， 如果把  O（n^3) 利用双指针把，复杂度降低到 O（n^2), 加强对于双指针的理解!
 */
func triple(a, b int) [3]int {
	c := -(a + b)

	if a > b {
		a, b = b, a
	}

	if c <= a {
		return [3]int{c, a, b}
	} else if c >= b {
		return [3]int{a, b, c}
	} else {
		return [3]int{a, c, b}
	}
}

func threeSum(nums []int) [][]int {
	hkeys := make(map[[3]int]struct{})
	ans := make([][]int, 0)
	hnums := make(map[int][]int, len(nums)) // keep the num appears index

	for i := 0; i < len(nums); i++ {
		if _, ok := hnums[nums[i]]; ok {
			hnums[nums[i]] = append(hnums[nums[i]], i)
		} else {
			hnums[nums[i]] = []int{i}
		}
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if v, exit := hnums[-(nums[i] + nums[j])]; exit {
				found := false
				// double check v
				for _, elem := range v {
					if elem != i && elem != j {
						found = true
						break
					}
				}
				if found {
					tri := triple(nums[i], nums[j])
					if _, ok := hkeys[tri]; !ok {
						ans = append(ans, []int{tri[0], tri[1], tri[2]})
						hkeys[tri] = struct{}{}
					}
				}
			}
		}
	}
	return ans
}
