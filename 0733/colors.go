package leetcode

func dfs(image [][]int, sr int, sc int, original int, newColor int) {
	m := len(image)
	n := len(image[0])
	if sr < 0 || sr >= m || sc < 0 || sc >= n {
		return
	}
	if image[sr][sc] == original {
		image[sr][sc] = newColor
		dfs(image, sr-1, sc, original, newColor)
		dfs(image, sr+1, sc, original, newColor)
		dfs(image, sr, sc-1, original, newColor)
		dfs(image, sr, sc+1, original, newColor)
	}
	return
}
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	if image[sr][sc] == newColor {
		return image
	}
	dfs(image, sr, sc, image[sr][sc], newColor)
	return image
}
