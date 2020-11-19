package tree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Build(nodes []*int) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: *nodes[0],
	}
	nodes = nodes[1:]
	level := []*TreeNode{root}
	for len(nodes) > 0 {
		nextLevel := make([]*TreeNode, 0)
		for _, node := range level {
			if len(nodes) > 0 {
				if nodes[0] != nil {
					left := &TreeNode{
						Val: *nodes[0],
					}
					node.Left = left
					nextLevel = append(nextLevel, left)
				}
				nodes = nodes[1:]
			}
			if len(nodes) > 0 {
				if nodes[0] != nil {
					right := &TreeNode{
						Val: *nodes[0],
					}
					node.Right = right
					nextLevel = append(nextLevel, right)
				}
				nodes = nodes[1:]
			}
			if len(nodes) == 0 {
				break
			}
		}
		level = nextLevel
	}
	return root
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

func PrintTreeNode(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	//初始化一个队列
	list := list.New()
	//从头部插入root
	list.PushFront(root)
	//开始层次遍历，在广度优先遍历基础上稍加调整
	for level := 0; list.Len() > 0; level++ {
		var currentLevel []int
		//取本层的节点数
		curentLenth := list.Len()
		for i := 0; i < curentLenth; i++ {
			//从尾部移除，Remove返回值为接口类型，需指定为TreeNode
			elem := list.Remove(list.Back())
			node := elem.(*TreeNode)
			if node == nil {
				currentLevel = append(currentLevel, 0x7fff)
				continue
			}
			//root默认从左往右
			currentLevel = append(currentLevel, node.Val)
			list.PushFront(node.Left)
			list.PushFront(node.Right)
		}
		//当前层结束
		result = append(result, currentLevel)
	}
	return result
}
