package main

func strToInt(s string) int {
	if s == "" {
		return 0
	}
	num := 0
	for _, v := range s {
		if v == '-' || v == '+' {
			continue
		}
		if v < 48 || v > 57 {
			return 0
		}
		num = num*10 + (int(v) - 48)
	}
	if s[0] == '-' {
		return -num
	}
	return num
}
