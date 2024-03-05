package main

import (
	"fmt"
)

func solution(my_string string, is_prefix string) int {
	if len(my_string) < len(is_prefix) {
		return 0
	}

	for i := 0; i < len(is_prefix); i++ {
		if my_string[i] != is_prefix[i] {
			return 0
		}
	}

	return 1
}
func main() {
	fmt.Println(solution("banana", "ban"))
	fmt.Println(solution("banana", "nan"))
	fmt.Println(solution("banana", "abcd"))
	fmt.Println(solution("banana", "bananan"))

}
