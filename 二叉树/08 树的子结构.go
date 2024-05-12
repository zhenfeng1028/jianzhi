package main

// 输入两棵二叉树A，B，判断B是不是A的子结构。

// 从树A的根节点开始寻找与B根节点相同的节点，如果相同，则继续比较其他节点，不相同，则继续。

type Node struct {
	val         int
	left, right *Node
}

func isSameTree(p, q *Node) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.val != q.val {
		return false
	}
	leftSame := isSameTree(p.left, q.left)
	rightSame := isSameTree(p.right, q.right)
	return leftSame && rightSame
}

func hasSubTree(p, q *Node) bool {
	result := false
	if p != nil && q != nil {
		if p.val == q.val {
			result = isSameTree(p, q)
			if !result {
				result = hasSubTree(p.left, q)
			}
		} else {
			result = hasSubTree(p.left, q)
		}
		if !result {
			result = hasSubTree(p.right, q)
		}
	}
	return result
}
