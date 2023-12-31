package main

import "fmt"

func solution(n int) int {
	n /= 2
	return n * (n + 1)
}

func main() {
	n := 10
	fmt.Println(solution(n))
}
