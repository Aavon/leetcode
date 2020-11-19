package rb

type NodeColor uint8

const Black NodeColor = 0
const Red NodeColor = 1

type RBTreeNode struct {
	Val    int
	Color  NodeColor
	Left   *RBTreeNode
	Right  *RBTreeNode
	Parent *RBTreeNode
}

func rotateLeft(node *RBTreeNode) {
	nextNode := node.Right
	node.Right = nextNode.Left

	if node.Parent != nil {
		if node.Parent.Left == node {
			node.Parent.Left = nextNode
		} else {
			node.Parent.Right = nextNode
		}
	}
	nextNode.Parent = node.Parent
	nextNode.Left = node
	node.Parent = nextNode
}

func rotateRight(node *RBTreeNode) {
	nextNode := node.Left
	node.Left = nextNode.Right

	if node.Parent != nil {
		if node.Parent.Left == node {
			node.Parent.Left = nextNode
		} else {
			node.Parent.Right = nextNode
		}
	}
	nextNode.Parent = node.Parent
	nextNode.Right = node
	node.Parent = nextNode
}

func getUncle(node *RBTreeNode) *RBTreeNode {
	if node.Parent == nil {
		return nil
	}
	if node.Parent.Parent == nil {
		return nil
	}
	if node.Parent.Parent.Left == node.Parent {
		return node.Parent.Parent.Right
	} else {
		return node.Parent.Parent.Left
	}
}

func fixUp(node *RBTreeNode) {
	if node == nil || node.Parent == nil || node.Parent.Color == Black {
		return
	}
	uncle := getUncle(node)
	if uncle == nil || uncle.Color == Black {
		node = node.Parent
		if node.Parent == node.Parent.Left {
			node.Color = Black
			if node.Parent.Parent != nil {
				node.Parent.Parent.Color = Red
				rotateRight(node.Parent.Parent)
			}
		} else {
			rotateLeft(node.Parent)
			fixUp(node)
		}
	} else {
		node.Parent.Color = Black
		uncle.Color = Black
		node.Parent.Parent = Red
		fixUp(node.Parent.Parent)
	}
}
