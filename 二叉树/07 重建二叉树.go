package main

// 输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
// 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。

// 根据前序遍历找到根节点，再根据中序遍历找到左右子树，由此得到左右子树的前序遍历和中序遍历，进行递归求解。

type Node struct {
	val         int
	left, right *Node
}

func reconstructTree(preOrder, inOrder []int) *Node {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}
	rootVal := preOrder[0]
	root := &Node{val: rootVal}
	// 寻找根节点在中序遍历中的位置
	i := -1
	for i = range inOrder {
		if inOrder[i] == rootVal {
			break
		}
	}
	leftPreOrder := preOrder[1 : i+1]
	leftInOrder := inOrder[:i]
	rightPreOrder := preOrder[i+1:]
	rightInOrder := inOrder[i+1:]
	leftChild := reconstructTree(leftPreOrder, leftInOrder)
	rightChild := reconstructTree(rightPreOrder, rightInOrder)
	root.left = leftChild
	root.right = rightChild
	return root
}
