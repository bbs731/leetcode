package leetcode

import "strconv"

/*
用 stack, 其它的花里胡哨的想法，都不对！
 */
func evalRPN(tokens []string) int {
	stack := []int{}

	for i := 0; i < len(tokens); i++ {
		num, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, num)
		} else {
			op1 := stack[len(stack)-1]
			op2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if tokens[i] == "+" {
				stack = append(stack, op1+op2)
			} else if tokens[i] == "-" {
				stack = append(stack, op2-op1)
			} else if tokens[i] == "*" {
				stack = append(stack, op2*op1)
			} else if tokens[i] == "/" {
				stack = append(stack, op2/op1)
			}
		}
	}
	return stack[0]
}
