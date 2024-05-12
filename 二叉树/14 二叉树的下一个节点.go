package main

// 给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针。

type Node struct {
	val         int
	left, right *Node
	next        *Node
}

func getNext(node *Node) *Node {
	if node == nil {
		return nil
	}
	// 1.如果存在右子节点，则右子节点的最左节点为下一个节点;
	if node.right != nil {
		tmp := node.right
		for tmp.left != nil {
			tmp = tmp.left
		}
		return tmp
	}
	// 2.如果不存在右子节点，且该节点是父节点的左子节点，则父节点为下一个节点;
	if node.right == nil {
		if node.next != nil && node.next.left == node {
			return node.next
		}
	}
	// 3.如果不存在右子节点，且该节点是父节点的右子节点，则向上找到一个节点是该其父节点的左孩子，则该其父节点为下一个节点。
	if node.right == nil {
		if node.next != nil && node.next.right == node {
			tmp := node.next
			for tmp.next != nil {
				if tmp.next.left == tmp {
					return tmp.next
				}
				tmp = tmp.next
			}
		}
	}
	return nil
}
