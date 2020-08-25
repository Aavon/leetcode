package tree

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
