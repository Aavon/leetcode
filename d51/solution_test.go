package d51

import (
	"fmt"
	"testing"
)

var (
	colMarked = map[int]bool{}
	pie = map[int]bool{} 
	na = map[int]bool{}
	result = make([][]string, 0)
)

/**
col = col
pie: row = col + c --> row - col = c
na: row = c - col --> row + col = c

**/

func dfs(n, row int, cur []int) {
	if row >= n {
		result = append(result, gen(cur, n))
		return
	}

	for col:=0; col<n; col++ {
		// 判断当前是否满足
		if colMarked[col] || pie[row-col] || na[row+col] {
			continue
		}

		colMarked[col] = true
		pie[row-col] = true
		na[row+col] = true
		cur[row] = col 
		// 下一层遍历
		dfs(n, row + 1, cur)

		// 恢复标记
		colMarked[col] = false
		pie[row-col] = false
		na[row+col] = false
		cur[row] = 0
	}
}

func gen(result []int, n int) []string {
	single := make([]string, 0, n)
	for _, pos := range result {
		row := make([]byte, n)
		for i := range row {
			if i == pos {
				row[i] = 'Q'
			} else {
				row[i] = '.'
			}
		}
		single = append(single, string(row))
	}
	return single
}


func solveNQueens(n int) [][]string {
	colMarked = map[int]bool{}
	pie = map[int]bool{} 
	na = map[int]bool{}
	result = make([][]string, 0)

	dfs(n, 0, make([]int, n))
	return result
}

func Test_a(t *testing.T) {
	fmt.Println(solveNQueens(6))
}