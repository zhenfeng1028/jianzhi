package main

// 给定一个整数数组nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。
// 找出只出现一次的那两个元素。你可以按任意顺序返回答案。

func singleNumber(nums []int) []int {
	checked := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if checked[i] {
			continue
		}
		for j := 0; j != i && j < len(nums); j++ {
			if checked[j] {
				continue
			}
			if nums[i]^nums[j] == 0 {
				checked[i] = true
				checked[j] = true
				break
			}
		}
	}
	ret := []int{}
	for i, v := range checked {
		if !v {
			ret = append(ret, nums[i])
		}
	}
	return ret
}
