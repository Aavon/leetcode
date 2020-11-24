package d120

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	level := len(triangle)

	ops := make([][]int, level)
	var cmpMin int
	for i := level - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == level-1 {
				ops[i] = append(ops[i], triangle[i][j])
			} else {
				if ops[i+1][j] < ops[i+1][j+1] {
					cmpMin = ops[i+1][j]
				} else {
					cmpMin = ops[i+1][j+1]
				}
				ops[i] = append(ops[i], cmpMin+triangle[i][j])
			}
		}
	}

	return ops[0][0]
}

func minimumTotal2(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	level := len(triangle)
	var cmpMin int
	for i := level - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i < level-1 {
				if triangle[i+1][j] < triangle[i+1][j+1] {
					cmpMin = triangle[i+1][j]
				} else {
					cmpMin = triangle[i+1][j+1]
				}
				triangle[i][j] = cmpMin + triangle[i][j]
			}
		}
	}

	return triangle[0][0]
}

func Test_a(t *testing.T) {
	data := make([][]int, 0)
	data = append(data, []int{2})
	data = append(data, []int{3, 4})
	data = append(data, []int{6, 5, 7})
	data = append(data, []int{4, 1, 8, 3})
	fmt.Println(minimumTotal(data))
	fmt.Println(minimumTotal2(data))
}
