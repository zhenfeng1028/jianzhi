package main

// 在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
// 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

// 从右上角开始，若小于target，向下走，删除一行，若大于target，向左走，删除一列。

func findNumberIn2DArray(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	col := n - 1
	for i := 0; i < m; i++ {
		for j := col; j >= 0; j-- {
			if matrix[i][j] == target {
				return true
			} else if matrix[i][j] < target {
				break
			} else if matrix[i][j] > target {
				col--
			}
		}
	}
	return false
}
