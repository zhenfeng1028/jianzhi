package main

// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

type Node struct {
	val         int
	left, right *Node
}

func printFromTopToBottom(root *Node) []int {
	res := []int{}
	queue := []*Node{}
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		res = append(res, node.val)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
		queue = queue[1:]
	}
	return res
}
