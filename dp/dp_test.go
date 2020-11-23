package dp

import (
	"fmt"
	"testing"
)

func countThePath(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	opt := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		opt[i] = make([]int, len(grid[i]))
	}
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if i+1 >= len(grid) {
				opt[i][j] = 1
				continue
			}
			if j+1 >= len(grid[i]) {
				opt[i][j] = 1
				continue
			}
			if grid[i+1][j] == 0 {
				opt[i][j] += opt[i+1][j]
			}
			if grid[i][j+1] == 0 {
				opt[i][j] += opt[i][j+1]
			}
		}
	}
	fmt.Scanln()
	for _, r := range opt {
		fmt.Println(r)
	}
	return opt[0][0]
}

func Test_countThePath(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{1, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println(countThePath(grid))
}
