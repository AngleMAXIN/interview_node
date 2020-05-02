package leetcode

func numIslands(grid [][]byte) int {
	l := len(grid)
	if l == 0 {
		return 0
	}
	count := 0
	for row := 0; row < l; row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '1' {
				count++
				sink(grid, row, col)
			}
		}
	}
	return count
}

func sink(grid [][]byte, row, col int) {
	if row < 0 || row > len(grid)-1 || col < 0 || col > len(grid[0])-1 || grid[row][col] == '0' {
		return
	}
	grid[row][col] = '0'
	sink(grid, row-1, col)
	sink(grid, row+1, col)
	sink(grid, row, col+1)
	sink(grid, row, col-1)
}
