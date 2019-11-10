package d46

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (72.14%)
 * Likes:    428
 * Dislikes: 0
 * Total Accepted:    54.8K
 * Total Submissions: 75.7K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个没有重复数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */

func permute(nums []int) [][]int {
	rs := [][]int{}
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	selected := nums[0]
	subs := permute(nums[1:])
	for _, s := range subs {
		sl := len(s)
		for i:=0; i<=sl; i++ {
			r := make([]int, i, sl+1)
			copy(r, s[0:i])
			r = append(r, selected)
			r = append(r, s[i:]...)
			rs = append(rs, r)
		}
	}
	return rs
}

func Test_a(t *testing.T) {
	nums := []int{1,2,3}
	fmt.Println(permute(nums))
}
