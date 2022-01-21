package leetcode

/*
计算斜率， 用 float32 不是好方法，

化简成，没有最小公约数的形式，   y/x ， 然后用  [2]int{y, x} 作为 hash key 就可以
 */
func maxPoints(points [][]int) int {
	n := len(points)

	max := 0
	for i := 0; i < n; i++ {
		origin := points[i]
		ks := make(map[float32]int)
		vertical := 0
		for j := 0; j < n && j != i; j++ {
			if origin[0] == points[j][0] {
				vertical++
				if vertical > max {
					max = vertical
				}
			} else {
				k := float32(points[j][1]-origin[1]) / float32(points[j][0]-origin[0])
				ks[k]++
				if ks[k] > max {
					max = ks[k]
				}
			}
		}
	}
	return max
}
