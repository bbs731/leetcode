package leetcode

/*
用基本的二维的 Prefixsum  枚举所有的 x1, y1  O(n^2)  然后枚举所有的 x2, y2  (x2>=x1, y2>=y1） O(n^2)  一共是 O(n^4)我们回超时

还是根据 Prefix sum 的思想。
我们一行一行处理， 处理一行的时候, 譬如 r 行，相当于再处理  1 x Column 的矩阵， 没一行有 n^2 个这样的矩阵。
然后这些矩阵，往上看， 和之前的 [0... r-1] 行， column 宽度相同的，又组成了，前缀和。
这样我们有 n 行， 每一行保证处理的 problem 是 O(n^2) 个， 就能保证总体 O(n ^3) 的复杂度。
 */

/*
https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/solution/yuan-su-he-wei-mu-biao-zhi-de-zi-ju-zhen-8ym2/
官网题解： 典型的， 二维降维到一维， 然后解决一维的问题

以前，碰到过类似的问题
 */

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m := len(matrix)
	n := len(matrix[0])

	counts := 0

	// 每一行，我们有 O(n^2) 个 subproblem

	row_prefix_sum := make([]int, n+1)

	// 然后我们创建， 这 n^2 个 column 方向的前缀和数组

	col_prefix_sum := make([][]int, n*n+1)
	hashcoltables := make([]map[int]int, n*n+1)

	for i := 0; i <= n*n; i++ {
		hashcoltables[i] = make(map[int]int)
		hashcoltables[i][0] = 1
	}

	// 初始化 col_prefix_sum
	for i := 0; i <= n*n; i++ {
		col_prefix_sum[i] = make([]int, m+1)
	}

	for r := 0; r < m; r++ {
		// 既然不需要，计数 row 上面的 1xColumn 的矩阵， 那么我们就不需要保存 hashrow table 了。
		//hashrow := make(map[int]int) // for each row create a new hashrow map
		//hashrow[0] = 1
		row_prefix_sum[0] = 0
		for i := 1; i <= n; i++ {
			row_prefix_sum[i] = row_prefix_sum[i-1] + matrix[r][i-1]
			// row 也计数， col 也计数， 会计算重复
			//val := row_prefix_sum[i] - target
			//if v, ok := hashrow[val]; ok {
			//	counts += v
			//}
			//hashrow[row_prefix_sum[i]]++
		}
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				col_prefix_sum[i*n+j][r+1] = col_prefix_sum[i*n+j][r] + row_prefix_sum[j+1] - row_prefix_sum[i]
				val := col_prefix_sum[i*n+j][r+1]
				if v, ok := hashcoltables[i*n+j][val-target]; ok {
					counts += v
				}
				hashcoltables[i*n+j][val]++
			}
		}
	}
	return counts
}
