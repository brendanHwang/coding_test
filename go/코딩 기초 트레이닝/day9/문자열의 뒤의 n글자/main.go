package main

import "fmt"

func solution(my_string string, n int) string {
	return my_string[len(my_string)-n:]
}

func main() {
	my_string := "ProgrammerS123"
	n := 11

	fmt.Println(
		solution(my_string, n),
	)
}
