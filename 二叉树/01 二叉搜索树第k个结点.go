package main

type Node struct {
	val         int
	left, right *Node
}

// 二叉搜索树的中序遍历是有序的，利用中序遍历寻找

var count int

func kthNode(root *Node, k int) *Node {
	return inorder(root, k)
}

func inorder(node *Node, k int) *Node {
	if node != nil {
		leftNode := kthNode(node.left, k)
		if leftNode != nil {
			return leftNode
		}
		count++
		if count == k {
			return node
		}
		rightNode := kthNode(node.right, k)
		if rightNode != nil {
			return rightNode
		}
	}
	return nil
}
