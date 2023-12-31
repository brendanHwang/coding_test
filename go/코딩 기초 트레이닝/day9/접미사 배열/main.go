package main

import (
	"fmt"
	"sort"
)


func solution(my_string string) []string {
	my_strings := make([]string, 0)

	for i,_ := range my_string {
		my_strings = append(my_strings, string(my_string[i:]))
	}

	// 사전순 정렬
	sort.Strings(my_strings)


	return my_strings
}

func main() {
	my_string := "banana"
	fmt.Println(
		solution(my_string),
	)
}
