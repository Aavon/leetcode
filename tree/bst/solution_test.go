package bst

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 */

// @lc code=start

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrder(t *TreeNode) []int {
	preOrder := []int{}
	if t == nil {
		return preOrder
	}
	preOrder = append(preOrder, t.Val)
	left := PreOrder(t.Left)
	right := PreOrder(t.Right)
	preOrder = append(preOrder, left...)
	preOrder = append(preOrder, right...)
	return preOrder
}

func generateTrees(n int) []*TreeNode {
	return generateRoot(1, n)
}

func generateRoot(start, end int) []*TreeNode {
	result := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		var leftTrees, rightTrees []*TreeNode
		if i == start {
			leftTrees = []*TreeNode{nil}
		} else {
			leftTrees = generateRoot(start, i-1)
		}
		if i == end {
			rightTrees = []*TreeNode{nil}
		} else {
			rightTrees = generateRoot(i+1, end)
		}
		for l := 0; l < len(leftTrees); l++ {
			for r := 0; r < len(rightTrees); r++ {
				root := &TreeNode{
					Val:   i,
					Left:  leftTrees[l],
					Right: rightTrees[r],
				}
				result = append(result, root)
			}
		}
	}
	return result
}

func numTrees(n int) int {
	subResult := make([]int, n+1)
	subResult[0] = 1
	subResult[1] = 1
	// 从n=2，往后计算
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			subResult[i] += subResult[j-1] * subResult[i-j]
		}
	}
	return subResult[n]
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	maxLeft := maxNode(root.Left)
	minRight := minNode(root.Right)
	var validLeft, validRight bool
	validLeft = root.Left == nil || (maxLeft.Val < root.Val && isValidBST(root.Left))
	if validLeft {
		validRight = root.Right == nil || (minRight.Val > root.Val && isValidBST(root.Right))
	}
	return validLeft && validRight
}

var maxCache = map[*TreeNode]*TreeNode{}

func maxNode(root *TreeNode) *TreeNode {
	var max = root
	if max == nil {
		return max
	}
	if cached, ok := maxCache[max]; ok {
		return cached
	}
	maxLeft := maxNode(max.Left)
	maxRight := maxNode(max.Right)

	if maxLeft != nil && maxLeft.Val > max.Val {
		max = maxLeft
	}
	if maxRight != nil && maxRight.Val > max.Val {
		max = maxRight
	}
	maxCache[root] = max
	return max
}

var minCache = map[*TreeNode]*TreeNode{}

func minNode(root *TreeNode) *TreeNode {
	var min = root
	if min == nil {
		return min
	}
	if cached, ok := minCache[min]; ok {
		return cached
	}
	minLeft := minNode(min.Left)
	minRight := minNode(min.Right)

	if minLeft != nil && minLeft.Val < min.Val {
		min = minLeft
	}
	if minRight != nil && minRight.Val < min.Val {
		min = minRight
	}
	minCache[root] = min
	return min
}

func TestA(t *testing.T) {
	trees := generateTrees(3)
	for _, t := range trees {
		fmt.Println(PreOrder(t))
	}
	fmt.Println(len(trees))
}

func TestB(t *testing.T) {
	fmt.Println(numTrees(3))
}
