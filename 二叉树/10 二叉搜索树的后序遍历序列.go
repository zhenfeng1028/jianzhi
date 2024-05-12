package main

// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。
// 如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。

/*
	二叉搜索树的性质：
    所有左子树的节点小于根节点，所有右子树的节点值大于根节点的值。

	算法步骤：
	1、后序遍历的最后一个值为root，在前面的数组中找到第一个大于root值的位置。
    2、这个位置的前面是root的左子树，右边是右子树。然后左右子树分别进行这个递归操作。
    3、其中，如果右边子树中有比root值小的直接返回false
*/

func verifyPostorder(postorder []int) bool {
	n := len(postorder)
	if n == 0 || n == 1 {
		return true
	}
	root := postorder[n-1]
	rightChildStartIndex := 0
	for i, v := range postorder {
		if v > root {
			rightChildStartIndex = i
			break
		}
	}
	leftChild := postorder[:rightChildStartIndex]
	rightChild := postorder[rightChildStartIndex : n-1]
	for _, v := range rightChild {
		if v < root {
			return false
		}
	}
	return verifyPostorder(leftChild) && verifyPostorder(rightChild)
}
