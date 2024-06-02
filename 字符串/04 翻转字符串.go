package main

import "strings"

func reverseSentence(s string) string {
	if s == "" {
		return ""
	}
	strs := strings.Split(s, " ")
	length := len(strs)
	if length == 1 {
		return strs[0]
	}
	orderedStrs := make([]string, length)
	for i, str := range strs {
		orderedStrs[length-1-i] = str
	}
	return strings.Join(orderedStrs, " ")
}
