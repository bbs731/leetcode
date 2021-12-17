package leetcode

func TwoSum(nums []int, target int) []int {

	m := make(map[int]int, len(nums))
	for i, n := range nums {
		m[n] = i
	}

	for i, n := range nums {
		if j, ok := m[target-n]; ok && i != j {
			// i ！= j  是个特殊情况， 譬如， target=6， num[2] = 3, 不应该返回 [2,2]
			return []int{i, j}
		}
	}
	return nil
}

// 更佳简洁的写法

func TwoSum1(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{j, i}
		}
		m[n] = i
	}
	return nil
}
