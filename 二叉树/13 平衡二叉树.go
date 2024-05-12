package main

// 给定一个二叉树，判断它是否是平衡二叉树
// 平衡二叉树是指该树所有节点的左右子树的深度相差不超过1。

type Node struct {
	val         int
	left, right *Node
}

func isBalanced(root *Node) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.left) && isBalanced(root.right) && abs(height(root.left)-height(root.right)) <= 1
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return max(height(node.left), height(node.right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
