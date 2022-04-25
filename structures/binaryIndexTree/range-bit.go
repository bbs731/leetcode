package binaryIndexTree

type BIT struct {
	t1 []int // 保存的是 差分数组 bi  // index 是 从  1 到 n
	t2 []int // 保存的是  i*bi 差分数组
	n  int
}

func (b BIT) lowbit(x int) int {
	return x & (-x)
}

func (b BIT) add(x, v int) int {
	v1 := x * v
	for x <= b.n {
		b.t1[x] += v
		b.t2[x] += v1
		x += b.lowbit(x)
	}
}

func (b BIT) getsum(t []int, x int) int {
	ret := 0
	for x >= 1 {
		ret += t[x]
		x -= b.lowbit(x)
	}
	return ret
}

func (b BIT) add1(l, r int, v int) {
	b.add(l, v)
	b.add(r+1, -v)
}

func (b BIT) getsum1(l, r int) int {

	// sum(i ~r) bi *(r+1) -  sum(i ~r)bi*i
	return (r+1)*b.getsum(b.t1, r) - (l-1+1)*b.getsum(b.t1, l-1) - (b.getsum(b.t2, r) - b.getsum(b.t2, l-1))
}

/* 如何初始化？/
 */
