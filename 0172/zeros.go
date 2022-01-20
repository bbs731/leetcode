package leetcode

/*
我们来查查 n 中包含了多少个 5.
然后你注意到 25 是 5 个 5， 有多包含了一个 5，  因此， 计算 n 时候 如果 n = x *5  然后我们还要看 x 又包含了多少个 5

例子 ： n =99
99 = 19 * 5 + 4
19 = 3*5 + 4
所以 99 一共有 19 + 3 个 5 也就是 22个 5， 因此有 22个 0
*/

/*
这他妈的就是在考察智商！
 */
func trailingZeroes(n int) int {

	ans := n / 5
	r := n / 5
	for r != 0 {
		ans += r / 5
		r = r / 5
	}
	return ans
}
