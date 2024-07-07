package main

// 例如序列[9,10,11,12,13,14,15,16]和序列[18,19,20,21,22]的和均为100

// 初始化small=1，big=2；
// small到big序列和小于target，big++；大于target，small++；
// 当small增加到(1+target)/2时停止

func continuousSequences(target int) [][]int {
	small, big := 1, 2
	ret := [][]int{}
	for small < (1+target)/2 {
		if sum(small, big) < target {
			big++
		} else if sum(small, big) > target {
			small++
		} else {
			tmp := []int{}
			for i := small; i <= big; i++ {
				tmp = append(tmp, i)
			}
			ret = append(ret, tmp)
			small++
		}
	}
	return ret
}

func sum(small, big int) int {
	ret := 0
	for i := small; i <= big; i++ {
		ret += i
	}
	return ret
}
