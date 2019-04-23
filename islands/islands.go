package island

// leetcode is passing a byte string for w/e reason. Not bothering with type casting
const land, water = 49, 48

func numIslands(grid [][]byte) int {
	var count, rows, columns int

	if len(grid) == 0 {
		return 0
	}

	rows = len(grid)
	columns = len(grid[0])

	for i := range grid {

		for j := range grid[i] {
			if grid[i][j] == water {
				continue
			}

			search(grid, rows, columns, i, j)

			grid[i][j] = land
			count++
		}

	}

	return count
}

// search isn't returning values - its only job is to mark any land
// connected to the given cooridinate as visited.
func search(grid [][]byte, rows, columns, y, x int) {
	if grid[y][x] == water {
		return
	}

	grid[y][x] = water

	// can I move up?
	if y != 0 {
		search(grid, rows, columns, y-1, x)
	}

	// can I move down?
	if y != rows-1 {
		search(grid, rows, columns, y+1, x)
	}

	// can I move to the left?
	if x > 0 {
		search(grid, rows, columns, y, x-1)
	}

	// can I move to the right?
	if x < columns-1 {
		search(grid, rows, columns, y, x+1)
	}
}
