package main

// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

type Node struct {
	val         int
	left, right *Node
}

func printFromTopToBottom(root *Node) [][]int {
	res := [][]int{}
	queue := []*Node{}
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		n := len(queue) // n用于保存每一次队列的元素个数，即二叉树每一层节点的个数
		tmp := []int{}
		for len(tmp) < n {
			node := queue[0]
			tmp = append(tmp, node.val)
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
			queue = queue[1:]
		}
		res = append(res, tmp)
	}
	return res
}
