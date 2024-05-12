package main

// 输入一颗二叉树，输出源二叉树的镜像。

type Node struct {
	val         int
	left, right *Node
}

func imageTree(root *Node) *Node {
	image := &Node{val: root.val}
	if root.left != nil {
		image.right = imageTree(root.left)
	}
	if root.right != nil {
		image.left = imageTree(root.right)
	}
	return image
}
