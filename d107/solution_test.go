package d107

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 */

// @lc code=start
/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test_a(t *testing.T) {
	root := &TreeNode{
		Val: 3,
	}
	root.Left = &TreeNode{
		Val: 4,
	}
	root.Right = &TreeNode{
		Val: 5,
	}
	fmt.Println(levelOrderBottom(root))
}

func levelOrderBottom(root *TreeNode) [][]int {
	all := make([][]int, 0)
	if root == nil {
		return all
	}
	level := []*TreeNode{&TreeNode{Left: root}}
	for len(level) > 0 {
		vals := make([]int, 0)
		nextLevel := make([]*TreeNode, 0)
		for _, n := range level {
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
				vals = append(vals, n.Left.Val)
			}
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
				vals = append(vals, n.Right.Val)
			}
		}
		if len(vals) > 0 {
			all = append(all, vals)
		}
		level = nextLevel
	}
	l := len(all)
	for i := 0; i < l/2; i++ {
		all[i], all[l-i-1] = all[l-i-1], all[i]
	}
	return all
}
