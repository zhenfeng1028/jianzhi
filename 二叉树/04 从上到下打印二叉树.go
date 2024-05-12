package main

// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，
// 第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

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
	isOddLevel := false
	for len(queue) != 0 {
		n := len(queue) // n用于保存每一次队列的元素个数，即二叉树每一层节点的个数
		tmp := []int{}
		isOddLevel = !isOddLevel
		for len(tmp) < n {
			node := queue[0]
			if isOddLevel {
				tmp = append(tmp, node.val)
			} else {
				tmp = append([]int{node.val}, tmp...)
			}
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
