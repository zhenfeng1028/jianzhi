package main

type Node struct {
	val         int
	left, right *Node
}

func isSymmetrical(root *Node) bool {
	if root == nil {
		return true
	}
	return isCommon(root.left, root.right)
}

func isCommon(leftNode, rightNode *Node) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}
	if leftNode != nil && rightNode != nil {
		return leftNode.val == rightNode.val &&
			isCommon(leftNode.left, rightNode.right) &&
			isCommon(leftNode.right, rightNode.left)
	}
	return false
}
