package d52

import (
	"fmt"
	"testing"
)

var result int

func dfs(n, row, col, pie, na int) {
	if row >= n {
		result++
		return
	}
	bits := (^(col | pie | na)) & ((1 << uint(n)) - 1)
	for bits != 0 {
		p := bits & -bits
		dfs(
			n,
			row+1,
			col|p,
			(pie|p)<<1,
			(na|p)>>1,
		)
		bits = bits & (bits - 1)
	}
}

func totalNQueens(n int) int {
	result = 0
	dfs(n, 0, 0, 0, 0)
	return result
}

func Test_a(t *testing.T) {
	fmt.Println(totalNQueens(1))
}
