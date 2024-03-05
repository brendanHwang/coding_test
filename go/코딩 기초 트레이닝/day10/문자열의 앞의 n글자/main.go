package main

import "fmt"

func solution(my_string string, n int) string {
	return my_string[:n]
}

func main() {
	fmt.Println(solution("ProgrammerS123", 11))
	fmt.Println(solution("He110W0r1d", 5))

}
