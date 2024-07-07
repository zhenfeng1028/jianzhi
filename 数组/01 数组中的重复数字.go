package main

func findRepeatNumber(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		}
		m[v] = struct{}{}
	}
	return -1
}
