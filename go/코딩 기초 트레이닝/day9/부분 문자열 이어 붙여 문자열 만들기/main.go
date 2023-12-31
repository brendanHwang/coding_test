package main

import "fmt"

func solution(my_strings []string, parts [][]int) string {
	answer := ""
	for i, my_string := range my_strings {
		from, to := parts[i][0], parts[i][1]
		answer += my_string[from : to+1]
	}
	return answer
}

func main() {
	my_strings, parts := []string{"progressive", "hamburger", "hammer", "ahocorasick"}, [][]int{{0, 4}, {1, 2}, {3, 5}, {7, 7}}
	fmt.Println(solution(my_strings, parts))
}
