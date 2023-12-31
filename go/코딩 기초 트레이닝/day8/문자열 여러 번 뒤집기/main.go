package main

import (
	"fmt"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func reverseStr(s string, from int, to int) string {
	reversedSlice := Reverse(s[from : to+1])
	return s[:from] + string(reversedSlice) + s[to+1:] // 새로운 문자열 생성
}

func solution(my_string string, queries [][]int) string {
	for _, query := range queries {
		my_string = reverseStr(my_string, query[0], query[1])
	}
	return my_string
}

func main() {
    s := "rermgorpsam"
    rectangles := [][]int{{2, 3}, {0, 7}, {5, 9}, {6, 10}}
    fmt.Println(solution(s, rectangles))
}