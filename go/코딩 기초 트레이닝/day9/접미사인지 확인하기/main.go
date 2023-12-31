package main

import (
	"fmt"
	"strings"
)

func solution(my_string string, is_suffix string) int {

	if strings.HasSuffix(my_string, is_suffix) {
		return 1
	} else {
		return 0
	}

}

func main() {

	my_string := "banana"
	is_suffix := "wxyz"
	fmt.Println(solution(
		my_string, is_suffix,
	))
}
