package main

// 输入一棵二叉树的根节点，求该树的深度。
// 从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

type Node struct {
	val         int
	left, right *Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.left)
	rightDepth := maxDepth(root.right)
	return max(leftDepth, rightDepth) + 1
}
