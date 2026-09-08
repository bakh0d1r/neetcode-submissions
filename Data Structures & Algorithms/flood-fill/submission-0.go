func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	original := image[sr][sc]

	if original == color {
		return image
	}
	dfs(image, sr, sc, original, color)
	return image
}

func dfs(image [][]int, r, c, original, color int) {
	if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || image[r][c] != original {
		return
	}
	image[r][c] = color
	dfs(image, r-1, c, original, color) //up
	dfs(image, r+1, c, original, color) //down
	dfs(image, r, c+1, original, color) //right
	dfs(image, r, c-1, original, color) //left
}