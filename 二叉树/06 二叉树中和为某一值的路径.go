package main

// 输入一颗二叉树的根节点和一个整数，输出二叉树中结点值的和为输入整数的所有路径，返回的数组长度大的靠前。

type Node struct {
	val         int
	left, right *Node
}

var (
	paths [][]int
	tmp   []int
)

func findPath(root *Node, target int) {
	if root == nil {
		return
	}
	search(root, target)
}

func search(node *Node, target int) {
	tmp = append(tmp, node.val)
	if node.left == nil && node.right == nil {
		if node.val == target {
			path := []int{}
			path = append(path, tmp...)
			paths = append(paths, path) // 如果直接append tmp，最终得到的结果都是一样的
		}
	}
	if node.left != nil {
		search(node.left, target-node.val)
	}
	if node.right != nil {
		search(node.right, target-node.val)
	}
	if len(tmp) > 0 {
		tmp = tmp[:len(tmp)-1]
	}
}
