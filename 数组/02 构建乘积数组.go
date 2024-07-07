package main

// 给定一个数组A[0,1,…,n-1]，请构建一个数组B[0,1,…,n-1]，
// 其中B[i]的值是数组A中除了下标i以外的元素的积, 即B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。

func constructArr(arr []int) []int {
	length := len(arr)
	left := make([]int, length)
	right := make([]int, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			left[i] = 1
		} else {
			left[i] = arr[i-1] * left[i-1]
		}
	}
	for i := length - 1; i >= 0; i-- {
		if i == length-1 {
			right[i] = 1
		} else {
			right[i] = arr[i+1] * right[i+1]
		}
	}
	ret := []int{}
	for i := range arr {
		ret = append(ret, left[i]*right[i])
	}
	return ret
}
