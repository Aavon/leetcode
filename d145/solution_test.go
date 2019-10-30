package d145

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (68.10%)
 * Likes:    182
 * Dislikes: 0
 * Total Accepted:    35.5K
 * Total Submissions: 52K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 后序 遍历。
 *
 * 示例:
 *
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * 输出: [3,2,1]
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 *
 */

// @lc code=start

//Definition for a binary tree node.
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

// 递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)
	left = append(left, right...)
	left = append(left, root.Val)
	return left
}

// 非递归
func postorderTraversal2(root *TreeNode) []int {
	rs := []int{}
	if root == nil {
		return rs
	}
	stack := []*TreeNode{}
	visited := []bool{}
	sl := 0
	stack = append(stack, root)
	visited = append(visited, true)
	if root.Right != nil {
		stack = append(stack, root.Right)
		visited = append(visited, false)
	}
	if root.Left != nil {
		stack = append(stack, root.Left)
		visited = append(visited, false)
	}
	for sl = len(stack); sl> 0; sl = len(stack) {
		n := stack[sl-1]
		if visited[sl -1] || n.Left == nil && n.Right == nil {
			stack = stack[:sl-1]
			visited = visited[:sl-1]
			rs = append(rs, n.Val)
			continue
		}
		if n.Right != nil {
			stack = append(stack, n.Right)
			visited = append(visited, false)
		}
		if n.Left != nil {
			stack = append(stack, n.Left)
			visited = append(visited, false)
		}
		visited[sl-1] = true
	}
	return rs
}


func Test_a(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val:2}
	root.Right.Left = &TreeNode{Val:3}
	rl := postorderTraversal2(root)
	fmt.Println(rl)
}