package d543

import (
	"encoding/json"
	"fmt"
	"leetcode/tree"
	"testing"
)

/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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

type TreeNode = tree.TreeNode

func Test_a(t *testing.T) {
	data := make([]*int, 0)
	err := json.Unmarshal([]byte("[4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]"), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	root := tree.Build(data)

	fmt.Println(tree.PreOrder(root))

	fmt.Println(tree.PrintTreeNode(root))

	fmt.Println(diameterOfBinaryTree(root))
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, maxD := helper(root)
	return maxD
}

func helper(root *TreeNode) (maxDepth, maxD int) {
	var left, ld int
	var right, rd int
	if root == nil {
		return
	}
	if root.Left != nil {
		left, ld = helper(root.Left)
		left += 1
	}
	if root.Right != nil {
		right, rd = helper(root.Right)
		right += 1
	}
	if left > right {
		maxDepth = left
	} else {
		maxDepth = right
	}
	maxD = left + right
	if ld > maxD {
		maxD = ld
	}
	if rd > maxD {
		maxD = rd
	}
	return
}
