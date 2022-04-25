package leetcode

/*
oh my god! you do it in 2D dimensions

二维的 BIT 可以用 [][]int 直接作为基础数据， 也可以像本题一样  []BIT
 */
type BIT struct {
	m    int
	tree []int // index 从 1 开始
}

func lowbit(x int) int {
	return x & (-x)
}

func (b *BIT) add(x int, k int) {
	for x <= b.m {
		b.tree[x] += k
		x += lowbit(x)
	}
}

func (b *BIT) query(x int) int {
	ans := 0
	for x >= 1 {
		ans += b.tree[x]
		x -= lowbit(x)
	}
	return ans
}

type NumMatrix struct {
	n      int
	C      []BIT // index 从 1 开始
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])

	C := make([]BIT, n+1)
	for i := 0; i < n+1; i++ {
		C[i].m = m
		C[i].tree = make([]int, m+1)
	}

	/*
	初始化 C
	 */
	for i := 0; i < n; i++ {
		col := i + 1
		for col <= n {
			for j := 0; j < m; j++ {
				C[col].add(j+1, matrix[j][i])
			}
			col += lowbit(col)
		}
	}
	return NumMatrix{
		n,
		C,
		matrix,
	}
}

func (this *NumMatrix) Update(row int, col int, val int) {
	origin := this.matrix[row][col]
	this.matrix[row][col] = val
	col += 1
	for col <= this.n {
		this.C[col].add(row+1, val-origin)
		col += lowbit(col)
	}
}

func (this *NumMatrix) Query(row int, col int) int {
	if row < 0 || col < 0 {
		return 0
	}
	ans := 0
	col = col + 1
	for col >= 1 {
		ans += this.C[col].query(row + 1)
		col -= lowbit(col)
	}
	return ans
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := 0

	ans += this.Query(row2, col2)
	ans -= this.Query(row2, col1-1)
	ans -= this.Query(row1-1, col2)
	ans += this.Query(row1-1, col1-1)
	return ans
}
