package s4

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Aavon/leetcode/tree"
)

// 二叉树的Z字遍历

func solution(root *tree.TreeNode) []int {
	var level int
	var leftStack []*tree.TreeNode
	var rightQueue []*tree.TreeNode
	result := make([]int, 0)
	rightQueue = append(rightQueue, root)

	for len(leftStack) > 0 || len(rightQueue) > 0 {
		if level%2 > 0 {
			// left
			for len(leftStack) > 0 {
				node := leftStack[len(leftStack)-1]
				leftStack = leftStack[:len(leftStack)-1]
				result = append(result, node.Val)
				if node.Right != nil {
					rightQueue = append(rightQueue, node.Right)
				}
				if node.Left != nil {
					rightQueue = append(rightQueue, node.Left)
				}
			}
		} else {
			// right
			for len(rightQueue) > 0 {
				node := rightQueue[len(rightQueue)-1]
				rightQueue = rightQueue[:len(rightQueue)-1]
				result = append(result, node.Val)
				if node.Left != nil {
					leftStack = append(leftStack, node.Left)
				}
				if node.Right != nil {
					leftStack = append(leftStack, node.Right)
				}
			}
		}
		level++
	}
	return result
}

func Test_a(t *testing.T) {
	var data []*int
	json.Unmarshal([]byte("[1,2,3,4,5,6,null,8,9,10,11]"), &data)
	root := tree.Build(data)
	vals := tree.PrintTreeNode(root)
	for _, v := range vals {
		fmt.Println(v)
	}

	fmt.Println(solution(root))
}
