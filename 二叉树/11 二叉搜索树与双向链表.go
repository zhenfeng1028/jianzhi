package main

// https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/description/

// 将一个二叉搜索树就地转化为一个已排序的双向循环链表 。
// 对于双向循环列表，你可以将左右孩子指针作为双向循环链表的前驱和后继指针，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。

/*
	1.排序链表：节点应从小到大排序，因此应使用中序遍历“从小到大”访问树的节点。
	2.双向链表：在构建相邻节点的引用关系时，设前驱节点pre和当前节点cur，不仅应构建pre.right = cur，也应构建cur.left = pre。
	3.循环链表：设链表头节点head和尾节点tail，则应构建head.left = tail和tail.right = head。
*/

type Node struct {
	val         int
	left, right *Node
}

var pre, head *Node

func treeToDoublyList(root *Node) *Node {
	inorder(root)
	head.left = pre
	pre.right = head
	return head
}

func inorder(cur *Node) {
	if cur == nil {
		return
	}
	inorder(cur.left)
	if pre != nil {
		pre.right = cur
	} else {
		head = cur
	}
	cur.left = pre
	pre = cur
	inorder(cur.right)
}
